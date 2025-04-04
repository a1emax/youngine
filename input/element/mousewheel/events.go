package mousewheel

import (
	"github.com/a1emax/youngine/basic"
)

// ScrollEvent occurs when wheel scroll is detected.
type ScrollEvent[B any] struct {
	Background B

	// Offset of wheel.
	Offset basic.Vec2
}

// Background of slave.
type Background[B any] struct {
	Background B

	// Offset of wheel.
	Offset basic.Vec2
}
