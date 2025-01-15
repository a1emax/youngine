package ebitenclock

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/clock"
)

// Second represents the number of ticks per second. This is the default value - use your own for custom cases.
const Second = ebiten.DefaultTPS * clock.Tick

// Minute represents the number of ticks per minute. This is the default value - use your own for custom cases.
const Minute = 60 * Second

// Hour represents the number of ticks per hour. This is the default value - use your own for custom cases.
const Hour = 60 * Minute
