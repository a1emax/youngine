package safeuint

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// Add returns x + y, or false if integer overflow occurred.
func Add[T basic.UnsignedInteger](x, y T) (T, bool) {
	if x > basic.MaxUint[T]()-y {
		return 0, false
	}

	return x + y, true
}

// MustAdd calls [Add] and panics if false is returned.
func MustAdd[T basic.UnsignedInteger](x, y T) T {
	result, ok := Add(x, y)
	if !ok {
		panic(fault.Trace(fault.ErrIntegerOverflow))
	}

	return result
}

// Sub returns x - y, or false if integer overflow occurred.
func Sub[T basic.UnsignedInteger](x, y T) (T, bool) {
	if y > x {
		return 0, false
	}

	return x - y, true
}

// MustSub calls [Sub] and panics if false is returned.
func MustSub[T basic.UnsignedInteger](x, y T) T {
	result, ok := Sub(x, y)
	if !ok {
		panic(fault.Trace(fault.ErrIntegerOverflow))
	}

	return result
}
