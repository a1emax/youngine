package basic

import (
	"math"
)

// Float is the basic floating-point type.
type Float = float64

// MaxFloat specifies the largest finite value representable by the [Float] type.
const MaxFloat = Float(math.MaxFloat64)

// SmallestNonzeroFloat specifies the smallest positive, non-zero value representable by the [Float] type.
const SmallestNonzeroFloat = Float(math.SmallestNonzeroFloat64)

// posInf specifies positive infinity.
var posInf = math.Inf(1)

// PosInf returns positive infinity.
func PosInf() Float {
	return posInf
}

// negInf specifies negative infinity.
var negInf = math.Inf(-1)

// NegInf returns negative infinity.
func NegInf() Float {
	return negInf
}
