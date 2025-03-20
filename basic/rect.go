package basic

import (
	"fmt"
	"image"
)

// Rect is rectangle with floating-point coordinates.
//
// By convention, right and bottom edges of rectangle are considered exclusive.
type Rect struct {
	Min, Size Vec2
}

// RectBtw returns rectangle with two opposite corners at points p0 and p1.
func RectBtw(p0, p1 Vec2) Rect {
	x0, y0 := p0.X(), p0.Y()
	x1, y1 := p1.X(), p1.Y()
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}

	return Rect{
		Min:  Vec2{x0, y0},
		Size: Vec2{x1 - x0, y1 - y0},
	}
}

// String implements the [fmt.Stringer] interface.
func (r Rect) String() string {
	return fmt.Sprintf("%s + %s", r.Min, r.Size)
}

// Image returns r converted to image-compatible format.
func (r Rect) Image() image.Rectangle {
	return image.Rectangle{
		Min: image.Pt(int(r.Left()), int(r.Top())),
		Max: image.Pt(int(r.Right()), int(r.Bottom())),
	}
}

// IsEmpty reports whether r contains no points.
func (r Rect) IsEmpty() bool {
	return r.Size.X() <= 0 || r.Size.Y() <= 0
}

// Left returns horizontal coordinate of left edge of r.
func (r Rect) Left() Float {
	return r.Min.X()
}

// Top returns vertical coordinate of top edge of r.
func (r Rect) Top() Float {
	return r.Min.Y()
}

// Right returns horizontal coordinate of right edge of r.
func (r Rect) Right() Float {
	return r.Min.X() + r.Size.X()
}

// Bottom returns vertical coordinate of bottom edge of r.
func (r Rect) Bottom() Float {
	return r.Min.Y() + r.Size.Y()
}

// Width returns horizontal size of r.
func (r Rect) Width() Float {
	return r.Size.X()
}

// Height returns vertical size of r.
func (r Rect) Height() Float {
	return r.Size.Y()
}

// Contains reports whether r contains point p.
func (r Rect) Contains(p Vec2) bool {
	return !r.IsEmpty() &&
		p.X() >= r.Left() && p.X() < r.Right() &&
		p.Y() >= r.Top() && p.Y() < r.Bottom()
}

// Inner returns r moved to origin.
func (r Rect) Inner() Rect {
	return Rect{
		Min:  Vec2{0, 0},
		Size: r.Size,
	}
}

// Move returns r moved by d.
func (r Rect) Move(d Vec2) Rect {
	return Rect{
		Min:  r.Min.Add(d),
		Size: r.Size,
	}
}

// MoveNum returns r moved by d in both dimensions.
func (r Rect) MoveNum(d Float) Rect {
	return r.Move(Vec2Num(d))
}

// Overlaps reports whether r has non-empty intersection with s.
func (r Rect) Overlaps(s Rect) bool {
	return !r.IsEmpty() && !s.IsEmpty() &&
		r.Left() < s.Right() && s.Left() < r.Right() &&
		r.Top() < s.Bottom() && s.Top() < r.Bottom()
}

// Resize returns r resized by d from center in both directions of both dimensions.
func (r Rect) Resize(d Vec2) Rect {
	return Rect{
		Min:  r.Min.Sub(d),
		Size: r.Size.Add(d.MulNum(2)),
	}
}

// ResizeNum returns r resized by d from center in all four directions.
func (r Rect) ResizeNum(d Float) Rect {
	return r.Resize(Vec2Num(d))
}

// Square returns square of r.
func (r Rect) Square() Float {
	return r.Size.X() * r.Size.Y()
}
