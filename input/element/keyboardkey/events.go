package keyboardkey

import (
	"github.com/a1emax/youngine/clock"
)

// DownEvent occurs when key press has just been detected.
type DownEvent[B any] struct {
	Background B

	// JustPressed reports whether key has just been pressed.
	JustPressed bool
}

// PressEvent occurs when key press is detected.
type PressEvent[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, that have passed since key press was detected.
	Duration clock.Ticks
}

// UpEvent occurs when key press has just been lost.
type UpEvent struct {

	// JustReleased bool reports whether key has just been released.
	JustReleased bool
}

// Background of slave.
type Background[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, that have passed since key press was detected.
	Duration clock.Ticks
}
