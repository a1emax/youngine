package ebiteninput

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/tempo"
)

// touchscreenTouchImpl is the implementation of the [input.TouchscreenTouch] interface.
type touchscreenTouchImpl struct {
	id        input.TouchscreenTouchID
	startedAt tempo.Time
	position  basic.Vec2
	isMarked  bool
}

// ID implements the [input.TouchscreenTouch] interface.
func (t *touchscreenTouchImpl) ID() input.TouchscreenTouchID {
	return t.id
}

// StartedAt implements the [input.TouchscreenTouch] interface.
func (t *touchscreenTouchImpl) StartedAt() tempo.Time {
	return t.startedAt
}

// Position implements the [input.TouchscreenTouch] interface.
func (t *touchscreenTouchImpl) Position() basic.Vec2 {
	return t.position
}

// IsMarked implements the [input.TouchscreenTouch] interface.
func (t *touchscreenTouchImpl) IsMarked() bool {
	return t.isMarked
}

// Mark implements the [input.TouchscreenTouch] interface.
func (t *touchscreenTouchImpl) Mark() {
	t.isMarked = true
}
