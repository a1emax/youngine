package mousewheel

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// Controller of wheel with background of type B.
type Controller[B any] interface {
	input.Controller[B]
}

// ControllerConfig configures [Controller].
type ControllerConfig[B any] struct {

	// Wheel state.
	Wheel input.MouseWheel

	// OnScroll, if specified, is called on [ScrollEvent].
	OnScroll func(event ScrollEvent[B])

	// Slave, if specified, is actuated if wheel is scrolled, or inhibited otherwise.
	Slave input.Controller[Background[B]]
}

// controllerImpl is the implementation of the [Controller] interface.
type controllerImpl[B any] struct {
	ControllerConfig[B]
}

// NewController initializes and returns new [Controller].
func NewController[B any](config ControllerConfig[B]) Controller[B] {
	if config.Wheel == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &controllerImpl[B]{
		ControllerConfig: config,
	}
}

// Actuate implements the [input.Controller] interface.
func (c *controllerImpl[B]) Actuate(background B) {
	if c.Wheel.IsMarked() {
		c.Inhibit()

		return
	}

	c.Wheel.Mark()

	wheelOffset := c.Wheel.Offset()
	if wheelOffset.IsZero() {
		c.Inhibit()

		return
	}

	if c.OnScroll != nil {
		c.OnScroll(ScrollEvent[B]{
			Background: background,
			Offset:     wheelOffset,
		})
	}

	if c.Slave != nil {
		c.Slave.Actuate(Background[B]{
			Background: background,
			Offset:     wheelOffset,
		})
	}
}

// Inhibit implements the [input.Controller] interface.
func (c *controllerImpl[B]) Inhibit() {
	if c.Slave != nil {
		c.Slave.Inhibit()
	}
}
