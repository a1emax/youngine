package mousecursor

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// Controller of cursor with background of type B.
type Controller[B any] interface {
	input.Controller[B]
}

// ControllerConfig configures [Controller].
type ControllerConfig[B any] struct {

	// Clock representing time.
	Clock clock.Clock

	// Input state.
	Input input.MouseCursor

	// HitTest, if specified, restricts active area cursor is being handled inside.
	HitTest func(position basic.Vec2) bool

	// OnEnter, if specified, is called on [EnterEvent].
	OnEnter func(event EnterEvent[B])

	// OnHover, if specified, is called on [HoverEvent].
	OnHover func(event HoverEvent[B])

	// OnLeave, if specified, is called on [LeaveEvent].
	OnLeave func(event LeaveEvent)

	// Slave, if specified, is actuated if cursor is inside active area, or inhibited otherwise.
	Slave input.Controller[Background[B]]
}

// controllerImpl is the implementation of the [Controller] interface.
type controllerImpl[B any] struct {
	ControllerConfig[B]

	detected   bool
	detectedAt clock.Time
}

// NewController initializes and returns new [Controller].
func NewController[B any](config ControllerConfig[B]) Controller[B] {
	if config.Input == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &controllerImpl[B]{
		ControllerConfig: config,
	}
}

// Actuate implements the [input.Controller] interface.
func (c *controllerImpl[B]) Actuate(background B) {
	if !c.Input.IsAvailable() || c.Input.IsMarked() {
		c.Inhibit()

		return
	}

	position := c.Input.Position()
	if c.HitTest != nil && !c.HitTest(position) {
		c.Inhibit()

		return
	}

	c.Input.Mark()

	now := c.Clock.Now()

	var duration clock.Ticks
	if c.detectedAt.IsZero() {
		c.detectedAt = now

		if c.OnEnter != nil {
			c.OnEnter(EnterEvent[B]{
				Background: background,
				Position:   position,
			})
		}
	} else {
		duration = now.Sub(c.detectedAt)
	}
	duration++ // Starts from 1.

	if c.OnHover != nil {
		c.OnHover(HoverEvent[B]{
			Background: background,
			Duration:   duration,
			Position:   position,
		})
	}

	if c.Slave != nil {
		c.Slave.Actuate(Background[B]{
			Background: background,
			Duration:   duration,
			Position:   position,
		})
	}
}

// Inhibit implements the [input.Controller] interface.
func (c *controllerImpl[B]) Inhibit() {
	if !c.detectedAt.IsZero() {
		if c.OnLeave != nil {
			c.OnLeave(LeaveEvent{})
		}

		c.detectedAt = clock.Time{}
	}

	if c.Slave != nil {
		c.Slave.Inhibit()
	}
}
