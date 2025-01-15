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

	// Button state.
	Button input.MouseButton

	// Clock representing time.
	Clock clock.Clock

	// OnDown, if specified, is called on [DownEvent].
	OnDown func(event DownEvent[B])

	// OnPress, if specified, is called on [PressEvent].
	OnPress func(event PressEvent[B])

	// OnUp, if specified, is called on [UpEvent].
	OnUp func(event UpEvent[B])

	// OnGone, if specified, is called on [GoneEvent].
	OnGone func(event GoneEvent)

	// Slave, if specified, is actuated if button is pressed, or inhibited otherwise.
	Slave input.Controller[Background[B]]
}

// controllerImpl is the implementation of the [Controller] interface.
type controllerImpl[B any] struct {
	ControllerConfig[B]

	buttonPressDetectedAt clock.Time
}

// NewController initializes and returns new [Controller].
func NewController[B any](config ControllerConfig[B]) Controller[B] {
	if config.Button == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if config.Clock == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &controllerImpl[B]{
		ControllerConfig: config,
	}
}

// Actuate implements the [input.Controller] interface.
func (c *controllerImpl[B]) Actuate(background B) {
	if c.Button.IsMarked() {
		c.Inhibit()

		return
	}

	c.Button.Mark()

	now := c.Clock.Now()

	if buttonPressedAt := c.Button.PressedAt(); !buttonPressedAt.IsZero() {
		var buttonPressDuration clock.Ticks
		if c.buttonPressDetectedAt.IsZero() {
			c.buttonPressDetectedAt = now

			if buttonPressedAt == now {
				if c.OnDown != nil {
					c.OnDown(DownEvent[B]{
						Background: background,
					})
				}
			}
		} else {
			buttonPressDuration = now.Sub(c.buttonPressDetectedAt)
		}
		buttonPressDuration++ // Starts from 1.

		if c.OnPress != nil {
			c.OnPress(PressEvent[B]{
				Background: background,
				Duration:   buttonPressDuration,
			})
		}

		if c.Slave != nil {
			c.Slave.Actuate(Background[B]{
				Background: background,
				Duration:   buttonPressDuration,
			})
		}
	} else {
		if c.Button.ReleasedAt() == now {
			if c.OnUp != nil {
				c.OnUp(UpEvent[B]{
					Background: background,
				})
			}
		}

		c.buttonPressDetectedAt = clock.Time{}

		if c.Slave != nil {
			c.Slave.Inhibit()
		}
	}
}

// Inhibit implements the [input.Controller] interface.
func (c *controllerImpl[B]) Inhibit() {
	if !c.buttonPressDetectedAt.IsZero() {
		if c.OnGone != nil {
			c.OnGone(GoneEvent{})
		}
	}

	c.buttonPressDetectedAt = clock.Time{}

	if c.Slave != nil {
		c.Slave.Inhibit()
	}
}
