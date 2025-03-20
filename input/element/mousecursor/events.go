package mousecursor

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/clock"
)

// EnterEvent occurs when cursor has just been detected.
type EnterEvent[B any] struct {
	Background B

	// Position of cursor.
	Position basic.Vec2
}

// HoverEvent occurs when cursor is detected.
type HoverEvent[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, that have passed since cursor was detected.
	Duration clock.Ticks

	// Position of cursor.
	Position basic.Vec2
}

// LeaveEvent occurs when cursor has just been lost.
type LeaveEvent struct {
}

// Background of slave.
type Background[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, that have passed since cursor was detected.
	Duration clock.Ticks

	// Position of cursor.
	Position basic.Vec2
}
