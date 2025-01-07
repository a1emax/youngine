package input

import (
	"github.com/a1emax/youngine/basic"
)

// MouseCursor state.
type MouseCursor interface {

	// IsAvailable reports whether cursor is available.
	IsAvailable() bool

	// Position returns position of cursor if it is available, or panics otherwise.
	Position() basic.Vec2

	// IsMarked reports whether state is marked. Marked state should not be handled.
	IsMarked() bool

	// Mark marks state until the next update.
	Mark()
}
