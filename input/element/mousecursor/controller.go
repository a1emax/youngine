package mousecursor

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// Controller of cursor with background of type B.
type Controller[B any] interface {
	input.Controller[B]
}

// ControllerConfig configures [Controller].
type ControllerConfig[B any] struct {

	// Cursor state.
	Cursor input.MouseCursor

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

	cursorDetected bool
}

// NewController initializes and returns new [Controller].
func NewController[B any](config ControllerConfig[B]) Controller[B] {
	if config.Cursor == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &controllerImpl[B]{
		ControllerConfig: config,
	}
}

// Actuate implements the [input.Controller] interface.
func (c *controllerImpl[B]) Actuate(background B) {
	if !c.Cursor.IsAvailable() || c.Cursor.IsMarked() {
		c.Inhibit()

		return
	}

	cursorPosition := c.Cursor.Position()
	if c.HitTest != nil && !c.HitTest(cursorPosition) {
		c.Inhibit()

		return
	}

	c.Cursor.Mark()

	if !c.cursorDetected {
		c.cursorDetected = true

		if c.OnEnter != nil {
			c.OnEnter(EnterEvent[B]{
				Background: background,
				Position:   cursorPosition,
			})
		}
	}

	if c.OnHover != nil {
		c.OnHover(HoverEvent[B]{
			Background: background,
			Position:   cursorPosition,
		})
	}

	if c.Slave != nil {
		c.Slave.Actuate(Background[B]{
			Background: background,
			Position:   cursorPosition,
		})
	}
}

// Inhibit implements the [input.Controller] interface.
func (c *controllerImpl[B]) Inhibit() {
	if c.cursorDetected {
		if c.OnLeave != nil {
			c.OnLeave(LeaveEvent{})
		}

		c.cursorDetected = false
	}

	if c.Slave != nil {
		c.Slave.Inhibit()
	}
}
