package promise

import (
	"github.com/a1emax/youngine/fault"
)

// When initializes and returns new promise that will be settled according to result of calling given callback
// when given thener is resolved. If thener is rejected, output promise will be rejected for the same reason
// without calling callback.
//
// Callback, in general case, should return new promise, along with settling which output promise will be settled.
// If nil is returned, output promise will be resolved immediately with zero resulting value of type T. If error
// is returned, output promise will be rejected immediately.
func When[T any](thener Thener, callback func() (Promise[T], error)) Promise[T] {
	if thener == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if callback == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	p, resolve, reject := New[T]()

	thener.Then(func() {
		callbackPromise, callbackErr := callback()

		if callbackErr != nil {
			reject(callbackErr)
		} else if callbackPromise != nil {
			callbackPromise.Then(func() {
				resolve(callbackPromise.Value())
			}, func(err error) {
				reject(err)
			})
		} else {
			var zero T
			resolve(zero)
		}
	}, func(err error) {
		reject(err)
	})

	return p
}

// Catch initializes and returns new promise that will be settled according to result of calling given callback
// when given thener is rejected. If thener is resolved, output promise will be rejected for default reason
// without calling callback.
//
// Callback, in general case, should return new promise, along with settling which output promise will be settled.
// If nil is returned, output promise will be resolved immediately with zero resulting value of type T. If error
// is returned, output promise will be rejected immediately.
func Catch[T any](thener Thener, callback func(err error) (Promise[T], error)) Promise[T] {
	if thener == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if callback == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	p, resolve, reject := New[T]()

	thener.Then(func() {
		reject(ErrRejected)
	}, func(err error) {
		callbackPromise, callbackErr := callback(err)

		if callbackErr != nil {
			reject(callbackErr)
		} else if callbackPromise != nil {
			callbackPromise.Then(func() {
				resolve(callbackPromise.Value())
			}, func(err error) {
				reject(err)
			})
		} else {
			var zero T
			resolve(zero)
		}
	})

	return p
}
