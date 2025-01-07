package colors

import (
	"fmt"

	"github.com/a1emax/youngine/basic"
)

// RGBA color with 8-bit channels.
type RGBA [4]uint8

// String implements the [fmt.Stringer] interface.
func (c RGBA) String() string {
	return fmt.Sprintf("0x%02X%02X%02X%02X", c[0], c[1], c[2], c[3])
}

// RGBA implements the [color.Color] interface.
func (c RGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(c[0])
	r |= r << 8
	g = uint32(c[1])
	g |= g << 8
	b = uint32(c[2])
	b |= b << 8
	a = uint32(c[3])
	a |= a << 8

	return r, g, b, a
}

// R returns red channel of c.
func (c RGBA) R() uint8 {
	return c[0]
}

// G returns green channel of c.
func (c RGBA) G() uint8 {
	return c[1]
}

// B returns blue channel of c.
func (c RGBA) B() uint8 {
	return c[2]
}

// A returns alpha channel of c.
func (c RGBA) A() uint8 {
	return c[3]
}

// Premul returns c premultiplied by alpha channel.
func (c RGBA) Premul() RGBA {
	f := basic.Float(c[3]) / 0xFF

	return RGBA{
		uint8(basic.Float(c[0]) * f),
		uint8(basic.Float(c[1]) * f),
		uint8(basic.Float(c[2]) * f),
		c[3],
	}
}

// ToFRGBA returns c converted to [FRGBA].
func (c RGBA) ToFRGBA() FRGBA {
	return FRGBA{
		basic.Float(c[0]) / 0xFF,
		basic.Float(c[1]) / 0xFF,
		basic.Float(c[2]) / 0xFF,
		basic.Float(c[3]) / 0xFF,
	}
}
