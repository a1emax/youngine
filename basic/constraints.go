package basic

// SignedInteger is constraint that permits any signed integer type.
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UnsignedInteger is constraint that permits any unsigned integer type.
type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is constraint that permits any integer type.
type Integer interface {
	SignedInteger | UnsignedInteger
}

// FloatingPoint is constraint that permits any floating-point type.
type FloatingPoint interface {
	~float32 | ~float64
}

// SignedNumeric is constraint that permits any signed numeric type.
type SignedNumeric interface {
	SignedInteger | FloatingPoint
}

// Numeric is constraint that permits any numeric type.
type Numeric interface {
	Integer | FloatingPoint
}

// Ordered is constraint that permits any type that supports operators < <= >= >.
type Ordered interface {
	Numeric | ~string
}
