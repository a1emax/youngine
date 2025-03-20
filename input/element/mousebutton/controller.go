package mousebutton

import (
	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// Controller of button with background of type B.
type Controller[B any] interface {
	input.Controller[B]
}

// ControllerConfig configures [Controller].
type ControllerConfig[B any] struct {

	// Clock representing time.
	Clock clock.Clock

	// Input state.
	Input input.MouseButton

	// OnDown, if specified, is called on [DownEvent].
	OnDown func(event DownEvent[B])

	// OnPress, if specified, is called on [PressEvent].
	OnPress func(event PressEvent[B])

	// OnUp, if specified, is called on [UpEvent].
	OnUp func(event UpEvent)

	// Slave, if specified, is actuated if button press is detected, or inhibited otherwise.
	Slave input.Controller[Background[B]]
}

// controllerImpl is the implementation of the [Controller] interface.
type controllerImpl[B any] struct {
	ControllerConfig[B]

	detectedAt clock.Time
}

// NewController initializes and returns new [Controller].
func NewController[B any](config ControllerConfig[B]) Controller[B] {
	if config.Clock == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if config.Input == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &controllerImpl[B]{
		ControllerConfig: config,
	}
}

// Actuate implements the [input.Controller] interface.
func (c *controllerImpl[B]) Actuate(background B) {
	if c.Input.IsMarked() {
		c.Inhibit()

		return
	}

	c.Input.Mark()

	now := c.Clock.Now()

	if pressedAt := c.Input.PressedAt(); !pressedAt.IsZero() {
		var duration clock.Ticks
		if c.detectedAt.IsZero() {
			c.detectedAt = now

			if c.OnDown != nil {
				c.OnDown(DownEvent[B]{
					Background:  background,
					JustPressed: pressedAt == now,
				})
			}
		} else {
			duration = now.Sub(c.detectedAt)
		}
		duration++ // Starts from 1.

		if c.OnPress != nil {
			c.OnPress(PressEvent[B]{
				Background: background,
				Duration:   duration,
			})
		}

		if c.Slave != nil {
			c.Slave.Actuate(Background[B]{
				Background: background,
				Duration:   duration,
			})
		}
	} else {
		if !c.detectedAt.IsZero() {
			if c.OnUp != nil {
				c.OnUp(UpEvent{
					JustReleased: c.Input.ReleasedAt() == now,
				})
			}

			c.detectedAt = clock.Time{}
		}

		if c.Slave != nil {
			c.Slave.Inhibit()
		}
	}
}

// Inhibit implements the [input.Controller] interface.
func (c *controllerImpl[B]) Inhibit() {
	if !c.detectedAt.IsZero() {
		if c.OnUp != nil {
			c.OnUp(UpEvent{
				JustReleased: false,
			})
		}

		c.detectedAt = clock.Time{}
	}

	if c.Slave != nil {
		c.Slave.Inhibit()
	}
}
