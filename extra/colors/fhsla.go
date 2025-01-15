package colors

import (
	"fmt"

	"github.com/a1emax/youngine/basic"
)

// FHSLA is HSLA color with floating-point channels.
type FHSLA [4]basic.Float

// String implements the [fmt.Stringer] interface.
func (c FHSLA) String() string {
	return fmt.Sprintf("hsla(%g, %g, %g, %g)", c[0], c[1], c[2], c[3])
}

// RGBA implements the [color.Color] interface.
func (c FHSLA) RGBA() (r, g, b, a uint32) {
	return c.ToFRGBA().RGBA()
}

// H returns hue channel of c.
func (c FHSLA) H() basic.Float {
	return c[0]
}

// S returns saturation channel of c.
func (c FHSLA) S() basic.Float {
	return c[1]
}

// L returns lightness channel of c.
func (c FHSLA) L() basic.Float {
	return c[2]
}

// A returns alpha channel of c.
func (c FHSLA) A() basic.Float {
	return c[3]
}

// Strict returns c with hue channel constrained to [0, 360] and all other ones constrained to [0, 1].
func (c FHSLA) Strict() FHSLA {
	return FHSLA{
		basic.Clamp(c[0], 0, 360),
		basic.Clamp(c[1], 0, 1),
		basic.Clamp(c[2], 0, 1),
		basic.Clamp(c[3], 0, 1),
	}
}

// Round returns c with all channels rounded to n decimal places.
func (c FHSLA) Round(n int) FHSLA {
	return FHSLA{
		basic.Round(c[0], n),
		basic.Round(c[1], n),
		basic.Round(c[2], n),
		basic.Round(c[3], n),
	}
}

// ToFRGBA returns c converted to [FRGBA].
func (c FHSLA) ToFRGBA() FRGBA {
	h := c[0]
	s := c[1]
	l := c[2]
	a := c[3]

	f := s * min(l, 1-l)

	k0 := basic.Mod(0+h/30, 12)
	k8 := basic.Mod(8+h/30, 12)
	k4 := basic.Mod(4+h/30, 12)

	r := l - f*max(-1, min(k0-3, min(9-k0, 1)))
	g := l - f*max(-1, min(k8-3, min(9-k8, 1)))
	b := l - f*max(-1, min(k4-3, min(9-k4, 1)))

	return FRGBA{r, g, b, a}
}
