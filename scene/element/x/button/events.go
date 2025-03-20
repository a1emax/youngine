package button

import (
	"github.com/a1emax/youngine/clock"
)

// ClickEvent occurs when button is clicked.
type ClickEvent struct {
}

// PressEvent occurs when button press is detected.
type PressEvent struct {
	Duration clock.Ticks
}
