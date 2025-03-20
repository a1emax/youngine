package mousebutton

import (
	"github.com/a1emax/youngine/clock"
)

// DownEvent occurs when button press has just been detected.
type DownEvent[B any] struct {
	Background B

	// JustPressed reports whether button has just been pressed.
	JustPressed bool
}

// PressEvent occurs when button press is detected.
type PressEvent[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, that have passed since button press was detected.
	Duration clock.Ticks
}

// UpEvent occurs when button press has just been lost.
type UpEvent struct {

	// JustReleased bool reports whether button has just been released.
	JustReleased bool
}

// Background of slave.
type Background[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, that have passed since button press was detected.
	Duration clock.Ticks
}
