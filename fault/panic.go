package fault

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// Recover calls given function and returns recovered error if there was panic, or nil otherwise.
func Recover(f func()) (err error) {
	defer func() {
		err = Recovered(recover())
	}()

	if f == nil {
		panic(Trace(ErrNilPointer))
	}

	f()

	return nil
}

// Recovered returns error represented by given panic recovery result, or nil if result is nil.
func Recovered(v any) error {
	if v == nil {
		return nil
	}

	err, ok := v.(error)
	if !ok {
		err = fmt.Errorf("panic recovered: %+v", v)
	}

	_, ok = err.(runtime.Error)
	if ok {
		err = fmt.Errorf("%w\n%s", err, debug.Stack())
	}

	return err
}
