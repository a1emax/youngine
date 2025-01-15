package input

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/clock"
)

// TouchscreenTouch state.
type TouchscreenTouch interface {

	// ID returns identifier of touch.
	ID() TouchscreenTouchID

	// StartedAt returns time of start of touch.
	StartedAt() clock.Time

	// Position returns position of touch.
	Position() basic.Vec2

	// IsMarked reports whether state is marked. Marked state should not be handled.
	IsMarked() bool

	// Mark marks state.
	Mark()
}
