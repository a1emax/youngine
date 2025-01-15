package colors

import (
	"fmt"
	"math"

	"github.com/a1emax/youngine/basic"
)

// FRGBA is RGBA color with floating-point channels.
type FRGBA [4]basic.Float

// String implements the [fmt.Stringer] interface.
func (c FRGBA) String() string {
	return fmt.Sprintf("rgba(%g, %g, %g, %g)", c[0], c[1], c[2], c[3])
}

// RGBA implements the [color.Color] interface.
func (c FRGBA) RGBA() (r, g, b, a uint32) {
	return c.ToRGBA().RGBA()
}

// R returns red channel of c.
func (c FRGBA) R() basic.Float {
	return c[0]
}

// G returns green channel of c.
func (c FRGBA) G() basic.Float {
	return c[1]
}

// B returns blue channel of c.
func (c FRGBA) B() basic.Float {
	return c[2]
}

// A returns alpha channel of c.
func (c FRGBA) A() basic.Float {
	return c[3]
}

// Strict returns c with all channels constrained to [0, 1].
func (c FRGBA) Strict() FRGBA {
	return FRGBA{
		basic.Clamp(c[0], 0, 1),
		basic.Clamp(c[1], 0, 1),
		basic.Clamp(c[2], 0, 1),
		basic.Clamp(c[3], 0, 1),
	}
}

// Premul returns c premultiplied by alpha channel.
func (c FRGBA) Premul() FRGBA {
	return FRGBA{
		c[0] * c[3],
		c[1] * c[3],
		c[2] * c[3],
		c[3],
	}
}

// Round returns c with all channels rounded to n decimal places.
func (c FRGBA) Round(n int) FRGBA {
	return FRGBA{
		basic.Round(c[0], n),
		basic.Round(c[1], n),
		basic.Round(c[2], n),
		basic.Round(c[3], n),
	}
}

// ToRGBA returns c converted to [RGBA].
func (c FRGBA) ToRGBA() RGBA {
	c = c.Strict()

	return RGBA{
		uint8(math.Round(c[0] * 0xFF)),
		uint8(math.Round(c[1] * 0xFF)),
		uint8(math.Round(c[2] * 0xFF)),
		uint8(math.Round(c[3] * 0xFF)),
	}
}
