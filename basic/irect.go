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

// Inner returns r moved to origin.
func (r Irect) Inner() Irect {
	return Irect{
		Min:  Ivec2{0, 0},
		Size: r.Size,
	}
}

// Contains reports whether r contains point p.
func (r Irect) Contains(p Ivec2) bool {
	return !r.IsEmpty() &&
		p.X() >= r.Left() && p.X() < r.Right() &&
		p.Y() >= r.Top() && p.Y() < r.Bottom()
}

// Overlaps reports whether r has non-empty intersection with s.
func (r Irect) Overlaps(s Irect) bool {
	return !r.IsEmpty() && !s.IsEmpty() &&
		r.Left() < s.Right() && s.Left() < r.Right() &&
		r.Top() < s.Bottom() && s.Top() < r.Bottom()
}
