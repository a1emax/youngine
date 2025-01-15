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

// Inner returns r moved to origin.
func (r Rect) Inner() Rect {
	return Rect{
		Min:  Vec2{0, 0},
		Size: r.Size,
	}
}

// Contains reports whether r contains point p.
func (r Rect) Contains(p Vec2) bool {
	return !r.IsEmpty() &&
		p.X() >= r.Left() && p.X() < r.Right() &&
		p.Y() >= r.Top() && p.Y() < r.Bottom()
}

// Overlaps reports whether r has non-empty intersection with s.
func (r Rect) Overlaps(s Rect) bool {
	return !r.IsEmpty() && !s.IsEmpty() &&
		r.Left() < s.Right() && s.Left() < r.Right() &&
		r.Top() < s.Bottom() && s.Top() < r.Bottom()
}
