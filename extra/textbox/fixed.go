package textbox

import (
	"math"

	"golang.org/x/image/math/fixed"

	"github.com/a1emax/youngine/basic"
)

// Fixed is basic fixed-point type.
type Fixed = fixed.Int26_6

// FixedBase specifies the number of fractional bits of value of the [Fixed] type.
const FixedBase = 6

// FloatToFixed returns x converted from floating-point type T to the [Fixed] type.
func FloatToFixed[T basic.FloatingPoint](x T) Fixed {
	return Fixed(math.Round(float64(x) * (1 << FixedBase)))
}

// FixedToFloat returns x converted from the [Fixed] type to the [basic.Float] type.
func FixedToFloat(x Fixed) basic.Float {
	return basic.Float(x) / (1 << FixedBase)
}
