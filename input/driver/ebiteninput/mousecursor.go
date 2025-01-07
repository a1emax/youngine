package ebiteninput

import (
	"github.com/a1emax/youngine/input"
)

// MouseCursor state based on Ebitengine.
type MouseCursor interface {
	input.MouseCursor

	// Update updates state.
	Update()
}

// NewMouseCursor initializes and returns new [MouseCursor].
func NewMouseCursor() MouseCursor {
	return &mouseCursorImpl{}
}
