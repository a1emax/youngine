package keyboardkey

import (
	"github.com/a1emax/youngine/clock"
)

// DownEvent occurs when key has just been pressed.
type DownEvent[B any] struct {
	Background B
}

// PressEvent occurs when key is pressed.
type PressEvent[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, during which key is pressed.
	Duration clock.Ticks
}

// UpEvent occurs when key has just been released.
type UpEvent[B any] struct {
	Background B
}

// GoneEvent occurs when key has just gone.
type GoneEvent struct {
}

// Background of slave.
type Background[B any] struct {
	Background B

	// Duration specifies number of ticks, including current one, during which key is pressed.
	Duration clock.Ticks
}
