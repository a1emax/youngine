package basic

import (
	"math"
	"unsafe"
)

// Abs returns |x|.
func Abs[T SignedNumeric](x T) T {
	if x < 0 {
		return -x
	}

	return x
}

// Clamp returns x constrained to closed interval [low, high].
func Clamp[T Ordered](x, low, high T) T {
	return max(low, min(x, high))
}

// FloorPoz returns greatest integer value less than or equal to x if it is positive, or zero otherwise.
func FloorPoz[T FloatingPoint](x T) T {
	return T(math.Floor(float64(Poz(x))))
}

// Fract returns fractional path of x.
func Fract[T FloatingPoint](x T) T {
	a := math.Abs(float64(x))

	return T(a - math.Floor(a))
}

// IsMaxInt reports whether x is maximum value representable by type T.
func IsMaxInt[T SignedInteger](x T) bool {
	return x == MaxInt[T]()
}

// IsMaxUint reports whether x is maximum value representable by type T.
func IsMaxUint[T UnsignedInteger](x T) bool {
	return x == MaxUint[T]()
}

// IsMinInt reports whether x is minimum value representable by type T.
func IsMinInt[T SignedInteger](x T) bool {
	return x == MinInt[T]()
}

// Loop returns x, looped in closed interval [low, high].
func Loop[T SignedInteger](x, low, high T) T {
	if high <= low {
		return low
	}

	x -= low
	n := high - low + 1

	if x < 0 {
		return low + n - (-x-1)%n - 1
	}

	return low + x%n
}

// MaxInt returns maximum value representable by type T.
func MaxInt[T SignedInteger]() T {
	offset := unsafe.Sizeof(T(0))<<3 - 2

	return 1<<offset - 1 | 1<<offset
}

// MaxUint returns maximum value representable by type T.
func MaxUint[T UnsignedInteger]() T {
	return ^T(0)
}

// MinInt returns minimum value representable by type T.
func MinInt[T SignedInteger]() T {
	return -MaxInt[T]() - 1
}

// Mod returns x modulo y that is computed as x - y*floor(x/y).
func Mod[T Numeric](x, y T) T {
	return x - y*T(math.Floor(float64(x)/float64(y)))
}

// Poz returns x if it is positive, or zero otherwise.
func Poz[T SignedNumeric](x T) T {
	return max(0, x)
}

// Round returns x rounded to n decimal places.
func Round[T FloatingPoint](x T, n int) T {
	if n <= 0 {
		return T(math.Round(float64(x)))
	}

	f := math.Pow10(n)

	return T(math.Round(float64(x)*f) / f)
}

// Sign returns sign of x, i.e. 1 if x is positive, -1 if it is negative, or 0 if it is zero.
func Sign[T SignedNumeric](x T) T {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}

	return 0
}
