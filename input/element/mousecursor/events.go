package mousecursor

import (
	"github.com/a1emax/youngine/basic"
)

// EnterEvent occurs when cursor enters active area.
type EnterEvent[B any] struct {
	Background B

	// Position of cursor.
	Position basic.Vec2
}

// HoverEvent occurs when cursor is inside active area.
type HoverEvent[B any] struct {
	Background B

	// Position of cursor.
	Position basic.Vec2
}

// LeaveEvent occurs when cursor leaves active area.
type LeaveEvent struct {
}

// Background of input handled by slave.
type Background[B any] struct {
	Background B

	// Position of cursor.
	Position basic.Vec2
}
