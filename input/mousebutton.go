package input

import (
	"github.com/a1emax/youngine/clock"
)

// MouseButton state.
type MouseButton interface {

	// Code returns code of button.
	Code() MouseButtonCode

	// PressedAt returns time of start pressing button, if it is pressed, or zero otherwise.
	PressedAt() clock.Time

	// ReleasedAt returns time of the last release of button, if it was released, or zero otherwise.
	ReleasedAt() clock.Time

	// IsMarked reports whether state is marked. Marked state should not be handled.
	IsMarked() bool

	// Mark marks state until the next update.
	Mark()
}
