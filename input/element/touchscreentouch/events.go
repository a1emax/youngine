package touchscreentouch

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/tempo"
)

// StartEvent occurs when touch has just been started inside active area.
type StartEvent[B any] struct {
	Background B

	// Position of touch.
	Position basic.Vec2
}

// HoverEvent occurs when touch is inside active area.
type HoverEvent[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, during which touch is inside active area.
	Duration tempo.Ticks

	// Position of touch.
	Position basic.Vec2
}

// EndEvent occurs when touch has just been ended inside active area.
type EndEvent[B any] struct {
	Background B
}

// GoneEvent occurs when touch has just gone.
type GoneEvent struct {
}

// Background of input handled by slave.
type Background[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, during which touch is inside region.
	Duration tempo.Ticks

	// Position of touch.
	Position basic.Vec2
}
