package basic

import (
	"fmt"
	"image"
)

// Irect is rectangle with integer coordinates.
//
// By convention, right and bottom edges of rectangle are considered exclusive.
type Irect struct {
	Min, Size Ivec2
}

// IrectBtw returns rectangle with two opposite corners at points p0 and p1.
func IrectBtw(p0, p1 Ivec2) Irect {
	x0, y0 := p0.X(), p0.Y()
	x1, y1 := p1.X(), p1.Y()
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}

	return Irect{
		Min:  Ivec2{x0, y0},
		Size: Ivec2{x1 - x0, y1 - y0},
	}
}

// String implements the [fmt.Stringer] interface.
func (r Irect) String() string {
	return fmt.Sprintf("%s + %s", r.Min, r.Size)
}

// Image returns r converted to image-compatible format.
func (r Irect) Image() image.Rectangle {
	return image.Rectangle{
		Min: image.Pt(r.Left(), r.Top()),
		Max: image.Pt(r.Right(), r.Bottom()),
	}
}

// Prec returns precise version of r.
func (r Irect) Prec() Rect {
	return Rect{
		Min:  r.Min.Prec(),
		Size: r.Size.Prec(),
	}
}

// IsEmpty reports whether r contains no points.
func (r Irect) IsEmpty() bool {
	return r.Size.X() <= 0 || r.Size.Y() <= 0
}

// Left returns horizontal coordinate of left edge of r.
func (r Irect) Left() int {
	return r.Min.X()
}

// Top returns vertical coordinate of top edge of r.
func (r Irect) Top() int {
	return r.Min.Y()
}

// Right returns horizontal coordinate of right edge of r.
func (r Irect) Right() int {
	return r.Min.X() + r.Size.X()
}

// Bottom returns vertical coordinate of bottom edge of r.
func (r Irect) Bottom() int {
	return r.Min.Y() + r.Size.Y()
}

// Width returns horizontal size of r.
func (r Irect) Width() int {
	return r.Size.X()
}

// Height returns vertical size of r.
func (r Irect) Height() int {
	return r.Size.Y()
}

// Contains reports whether r contains point p.
func (r Irect) Contains(p Ivec2) bool {
	return !r.IsEmpty() &&
		p.X() >= r.Left() && p.X() < r.Right() &&
		p.Y() >= r.Top() && p.Y() < r.Bottom()
}

// Inner returns r moved to origin.
func (r Irect) Inner() Irect {
	return Irect{
		Min:  Ivec2{0, 0},
		Size: r.Size,
	}
}

// MoveNum returns r moved by d in both dimensions.
func (r Irect) MoveNum(d int) Irect {
	return r.Move(Ivec2Num(d))
}

// Move returns r moved by d.
func (r Irect) Move(d Ivec2) Irect {
	return Irect{
		Min:  r.Min.Add(d),
		Size: r.Size,
	}
}

// Overlaps reports whether r has non-empty intersection with s.
func (r Irect) Overlaps(s Irect) bool {
	return !r.IsEmpty() && !s.IsEmpty() &&
		r.Left() < s.Right() && s.Left() < r.Right() &&
		r.Top() < s.Bottom() && s.Top() < r.Bottom()
}

// ResizeNum returns r resized by d from center in all four directions.
func (r Irect) ResizeNum(d int) Irect {
	return r.Resize(Ivec2Num(d))
}

// Resize returns r resized by d from center in both directions of both dimensions.
func (r Irect) Resize(d Ivec2) Irect {
	return Irect{
		Min:  r.Min.Sub(d),
		Size: r.Size.Add(d.MulNum(2)),
	}
}

// Square returns square of r.
func (r Irect) Square() int {
	return r.Size.X() * r.Size.Y()
}
