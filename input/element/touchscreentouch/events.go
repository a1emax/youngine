package touchscreentouch

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/clock"
)

// DownEvent occurs when touch has just been detected.
type DownEvent[B any] struct {
	Background B

	// JustStarted reports whether touch has just been started.
	JustStarted bool

	// Position of touch.
	Position basic.Vec2
}

// HoverEvent occurs when touch is detected.
type HoverEvent[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, that have passed since touch was detected.
	Duration clock.Ticks

	// Position of touch.
	Position basic.Vec2
}

// UpEvent occurs when touch has just been lost.
type UpEvent struct {

	// JustEnded reports whether touch has just been ended (from controller's point of view).
	JustEnded bool
}

// Background of slave.
type Background[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, that have passed since touch was detected.
	Duration clock.Ticks

	// Position of touch.
	Position basic.Vec2
}
