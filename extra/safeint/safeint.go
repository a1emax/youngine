package safeint

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// Abs returns absolute value of x, or false if integer overflow occurred.
func Abs[T basic.SignedInteger](x T) (T, bool) {
	if x >= 0 {
		return x, true
	}

	maxT := basic.MaxInt[T]()

	if x == -maxT-1 {
		return 0, false
	}

	return -x, true
}

// MustAbs calls [Abs] and panics if false is returned.
func MustAbs[T basic.SignedInteger](x T) T {
	result, ok := Abs(x)
	if !ok {
		panic(fault.Trace(fault.ErrIntegerOverflow))
	}

	return result
}

// Neg returns x with inverted sign, or false if integer overflow occurred.
func Neg[T basic.SignedInteger](x T) (T, bool) {
	if x >= 0 {
		return -x, true
	}

	maxT := basic.MaxInt[T]()

	if x == -maxT-1 {
		return 0, false
	}

	return -x, true
}

// MustNeg calls [Neg] and panics if false is returned.
func MustNeg[T basic.SignedInteger](x T) T {
	result, ok := Neg(x)
	if !ok {
		panic(fault.Trace(fault.ErrIntegerOverflow))
	}

	return result
}

// Add returns sum of x and y, or false if integer overflow occurred.
func Add[T basic.SignedInteger](x, y T) (T, bool) {
	if x == 0 {
		return y, true
	}
	if y == 0 {
		return x, true
	}

	maxT := basic.MaxInt[T]()

	if x > 0 && y > 0 && x > maxT-y {
		return 0, false
	}

	if x < 0 && y < 0 && x < -maxT-1-y {
		return 0, false
	}

	return x + y, true
}

// MustAdd calls [Add] and panics if false is returned.
func MustAdd[T basic.SignedInteger](x, y T) T {
	result, ok := Add(x, y)
	if !ok {
		panic(fault.Trace(fault.ErrIntegerOverflow))
	}

	return result
}

// Sub returns difference of x and y, or false if integer overflow occurred.
func Sub[T basic.SignedInteger](x, y T) (T, bool) {
	if y == 0 {
		return x, true
	}

	maxT := basic.MaxInt[T]()

	if x == 0 {
		if y == -maxT-1 {
			return 0, false
		}

		return -y, true
	}

	if x > 0 && y < 0 && x > maxT+y {
		return 0, false
	}

	if x < 0 && y > 0 && x < -maxT-1+y {
		return 0, false
	}

	return x - y, true
}

// MustSub calls [Sub] and panics if false is returned.
func MustSub[T basic.SignedInteger](x, y T) T {
	result, ok := Sub(x, y)
	if !ok {
		panic(fault.Trace(fault.ErrIntegerOverflow))
	}

	return result
}
