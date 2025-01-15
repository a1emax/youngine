package input

import (
	"github.com/a1emax/youngine/basic"
)

// MouseWheel state.
type MouseWheel interface {

	// Offset returns offset of wheel if it is scrolled, or zero otherwise.
	Offset() basic.Vec2

	// IsMarked reports whether state is marked. Marked state should not be handled.
	IsMarked() bool

	// Mark marks state the next update.
	Mark()
}
