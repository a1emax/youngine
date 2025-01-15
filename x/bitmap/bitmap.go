package bitmap

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/x/colors"
)

// BPP specifies the number of BYTES per pixel.
const BPP = 4

// Bitmap is two-dimensional array of pixels.
//
// To use Bitmap as mask fill it with colors where all channels have the same value (i.e. non-premultiplied grayscale).
// To draw such mask use the [github.com/hajimehoshi/ebiten/v2.ColorScale.ScaleWithColor] method with premultiplied
// color - combination of this color and mask will be correctly premultiplied:
//
//	raw_c - source color's original component channel (red, green or blue)
//	src_c - source color's premultiplied component channel
//	src_a - source color's alpha channel
//	bmp   - bitmap color's channel (all the same)
//	dst_c - destination color's component channel
//	dst_a - destination color's alpha channel
//
//	src_c = raw_c * src_a
//	dst_a = src_a * bmp
//	dst_c = src_c * bmp = raw_c * src_a * bmp = raw_c * dst_a (premultiplied)
type Bitmap struct {
	data   []byte
	width  int
	height int
}

// New initializes and returns new [Bitmap] of given width and height.
func New(width, height int) Bitmap {
	if width < 0 {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}
	if height < 0 {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	return Bitmap{
		data:   make([]byte, BPP*width*height),
		width:  width,
		height: height,
	}
}

// IsNil reports whether bitmap is nil.
func (b Bitmap) IsNil() bool {
	return b.data == nil
}

// Width returns width of bitmap.
func (b Bitmap) Width() int {
	return b.width
}

// Height returns height of bitmap.
func (b Bitmap) Height() int {
	return b.height
}

// Data returns raw data of bitmap.
func (b Bitmap) Data() []byte {
	return b.data
}

// Get returns color of pixel at (x, y) if it is inside bitmap, or zero value otherwise.
func (b Bitmap) Get(x, y int) colors.RGBA {
	if x < 0 || x >= b.width || y < 0 || y >= b.height {
		return colors.RGBA{}
	}

	i := b.index(x, y)

	return colors.RGBA(b.data[i : i+BPP])
}

// Set sets pixel at (x, y), if it is inside bitmap, to given color, or does nothing otherwise.
func (b Bitmap) Set(x, y int, color colors.RGBA) {
	if x < 0 || x >= b.width || y < 0 || y >= b.height {
		return
	}

	i := b.index(x, y)

	copy(b.data[i:], color[:])
}

// Fill fills bitmap with given color.
func (b Bitmap) Fill(color colors.RGBA) {
	if b.width == 0 || b.height == 0 {
		return
	}

	copy(b.data, color[:])
	for i := BPP; i < len(b.data); i *= 2 {
		copy(b.data[i:], b.data[:i])
	}
}

// Line fills horizontal line from (x1, y) to (x2, y) with given color.
func (b Bitmap) Line(x1, x2, y int, color colors.RGBA) {
	if x1 > x2 || (x1 < 0 && x2 < 0) || x1 >= b.width || y < 0 || y >= b.height {
		return
	}

	x1 = basic.Clamp(x1, 0, b.width-1)
	x2 = basic.Clamp(x2, 0, b.width-1)

	i1 := b.index(x1, y)
	i2 := b.index(x2, y)

	b.line(i1, i2, color)
}

// Lines fills horizontal lines of given length from each (x, y) with given color.
func (b Bitmap) Lines(length int, xy [][2]int, color colors.RGBA) {
	if b.width == 0 || b.height == 0 || length <= 0 {
		return
	}

	var j1, j3 int
	dj := -1

	for _, p := range xy {
		x1 := p[0]
		x2 := x1 + length - 1
		y := p[1]
		if (x1 < 0 && x2 < 0) || x1 >= b.width || y < 0 || y >= b.height {
			continue
		}

		x1 = basic.Clamp(x1, 0, b.width-1)
		x2 = basic.Clamp(x2, 0, b.width-1)

		i1 := b.index(x1, y)
		i2 := b.index(x2, y)

		if di := i2 - i1; di <= dj {
			i3 := i2 + BPP
			copy(b.data[i1:i3], b.data[j1:j3])

			continue
		}

		j1 = i1
		j3, dj = b.line(i1, i2, color)
	}
}

// Rect fills rectangle from (x1, y1) to (x2, y2) with given color.
func (b Bitmap) Rect(x1, y1, x2, y2 int, color colors.RGBA) {
	if x1 > x2 || (x1 < 0 && x2 < 0) || x1 >= b.width || y1 > y2 || (y1 < 0 && y2 < 0) || y1 >= b.height {
		return
	}

	x1 = basic.Clamp(x1, 0, b.width-1)
	x2 = basic.Clamp(x2, 0, b.width-1)

	y1 = basic.Clamp(y1, 0, b.height-1)
	y2 = basic.Clamp(y2, 0, b.height-1)

	i1 := b.index(x1, y1)
	i2 := b.index(x2, y1)

	i3, _ := b.line(i1, i2, color)

	j1 := i1
	j3 := i3
	w := BPP * b.width
	for y := y1 + 1; y <= y2; y++ {
		j1 += w
		j3 += w
		copy(b.data[j1:j3], b.data[i1:i3])
	}
}

// index returns index of first byte of pixel at (x, y).
func (b Bitmap) index(x, y int) int {
	return BPP * (y*b.width + x)
}

// line fills horizontal line from i1 to i2 with given color and returns i2+BPP and i2-i1.
func (b Bitmap) line(i1, i2 int, color colors.RGBA) (int, int) {
	i3 := i2 + BPP
	di := i2 - i1

	copy(b.data[i1:], color[:])
	for ri := BPP; ri <= di; ri *= 2 {
		i := i1 + ri
		copy(b.data[i:i3], b.data[i1:i])
	}

	return i3, di
}
