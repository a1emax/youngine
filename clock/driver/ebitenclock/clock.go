package ebitenclock

import (
	"github.com/a1emax/youngine/clock"
)

// Clock of logical time based on Ebitengine.
type Clock interface {
	clock.Clock

	// Update updates clock.
	Update()
}

// clockImpl is the implementation of the [Clock] interface.
type clockImpl struct {
	ticks clock.Ticks
}

// New initializes and returns new [Clock].
func New() Clock {
	return &clockImpl{}
}

// Now implements the [clock.Clock] interface.
func (c *clockImpl) Now() clock.Time {
	return clock.At(c.ticks)
}

// Update implements the [Clock] interface.
func (c *clockImpl) Update() {
	c.ticks++
}
