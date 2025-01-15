package mousewheel

import (
	"github.com/a1emax/youngine/basic"
)

// ScrollEvent occurs when wheel is scrolled.
type ScrollEvent[B any] struct {
	Background B

	// Offset of wheel.
	Offset basic.Vec2
}

// Background of input handled by slave.
type Background[B any] struct {
	Background B

	// Offset of wheel.
	Offset basic.Vec2
}
