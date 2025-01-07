package scope

import (
	"errors"

	"github.com/a1emax/youngine/fault"
)

// Setup sets up new scope.
func Setup(init InitFunc) (TeardownFunc, error) {
	if init == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return func() (_ TeardownFunc, _err error) {
		var _ok bool
		var deferred []func()

		var afterTeardownStart bool
		teardown := func() {
			if afterTeardownStart {
				panic(fault.Trace(fault.ErrInvalidUse))
			}
			afterTeardownStart = true

			defer func() {
				deferred = nil
			}()

			var err error
			for i := len(deferred) - 1; i >= 0; i-- {
				err = errors.Join(err, fault.Recover(deferred[i]))
			}
			if err != nil {
				panic(err)
			}
		}

		defer func() {
			if _ok {
				return
			}

			teardownErr := fault.Recover(teardown)
			if teardownErr == nil {
				return
			}

			var setupErr error
			if _err == nil {
				setupErr = fault.Recovered(recover())
			} else {
				setupErr = _err
				_err = nil
			}

			panic(errors.Join(setupErr, teardownErr))
		}()

		var afterInitEnd bool
		defer func() {
			afterInitEnd = true
		}()
		lc := &lifecycleImpl{
			_defer: func(f func()) {
				if f == nil {
					panic(fault.ErrNilPointer)
				}
				if afterInitEnd {
					panic(fault.Trace(fault.ErrInvalidUse))
				}

				deferred = append(deferred, f)
			},
		}

		err := init(lc)
		if err != nil {
			return nil, err
		}

		_ok = true

		return teardown, nil
	}()
}

// MustSetup is like [Setup], but assumes panicking instead of returning error.
func MustSetup(init MustInitFunc) TeardownFunc {
	if init == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	teardown, err := Setup(func(lifetime Lifecycle) error {
		init(lifetime)

		return nil
	})
	if err != nil {
		panic(err)
	}

	return teardown
}

// InitFunc initializes scope.
//
// If InitFunc panics or returns error, deferred functions are called internally in reverse order.
type InitFunc func(lc Lifecycle) error

// MustInitFunc is like [InitFunc], but assumes panicking instead of returning error.
type MustInitFunc func(lc Lifecycle)

// Lifecycle of scope.
type Lifecycle interface {

	// Defer defers given function call during scope initialization until scope teardown.
	Defer(f func())
}

// lifecycleImpl is the implementation of the [Lifecycle] interface.
type lifecycleImpl struct {
	_defer func(f func())
}

// Defer implements the [Lifecycle] interface.
func (l *lifecycleImpl) Defer(f func()) {
	l._defer(f)
}

// TeardownFunc tears down scope by calling deferred functions in reverse order.
//
// TeardownFunc panics if called more than once.
type TeardownFunc func()
