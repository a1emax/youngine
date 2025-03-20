package touchscreentouch

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// Controller of touch with background of type B.
type Controller[B any] interface {
	input.Controller[B]
}

// ControllerConfig configures [Controller].
type ControllerConfig[B any] struct {

	// Clock representing time.
	Clock clock.Clock

	// Input state.
	Input input.Touchscreen

	// HitTest, if specified, restricts active area touch is being handled inside.
	HitTest func(position basic.Vec2) bool

	// OnDown, if specified, is called on [DownEvent].
	OnDown func(event DownEvent[B])

	// OnHover, if specified, is called on [HoverEvent].
	OnHover func(event HoverEvent[B])

	// OnUp, if specified, is called on [UpEvent].
	OnUp func(event UpEvent)

	// Slave, if specified, is actuated if touch is detected, or inhibited otherwise.
	Slave input.Controller[Background[B]]
}

// controllerImpl is the implementation of the [Controller] interface.
type controllerImpl[B any] struct {
	ControllerConfig[B]

	detectedAt clock.Time
	id         input.TouchscreenTouchID // Valid only if detectedAt is not zero.
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
	var detectedFound bool
	var touch input.TouchscreenTouch
	var id input.TouchscreenTouchID
	var position basic.Vec2
	for i, n := 0, c.Input.TouchCount(); i < n; i++ {
		currentTouch := c.Input.Touch(i)

		currentID := currentTouch.ID()
		if !c.detectedAt.IsZero() && currentID == c.id {
			detectedFound = true
		}

		if currentTouch.IsMarked() {
			continue
		}

		currentPosition := currentTouch.Position()
		if c.HitTest != nil && !c.HitTest(currentPosition) {
			continue
		}

		// Check for multiple touches.
		if touch != nil {
			touch = nil

			break
		}

		touch = currentTouch
		id = currentID
		position = currentPosition
	}

	up := func() {
		if c.OnUp != nil {
			c.OnUp(UpEvent{
				JustEnded: !detectedFound,
			})
		}

		c.detectedAt = clock.Time{}
		c.id = 0
	}

	if touch == nil {
		if !c.detectedAt.IsZero() {
			up()
		}

		if c.Slave != nil {
			c.Slave.Inhibit()
		}

		return
	}

	touch.Mark()

	// Check for different touch.
	if !c.detectedAt.IsZero() && id != c.id {
		up()
	}

	now := c.Clock.Now()

	var duration clock.Ticks
	if c.detectedAt.IsZero() {
		c.detectedAt = now
		c.id = id

		if c.OnDown != nil {
			c.OnDown(DownEvent[B]{
				Background:  background,
				JustStarted: touch.StartedAt() == now,
				Position:    position,
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
		if c.OnUp != nil {
			c.OnUp(UpEvent{
				JustEnded: false,
			})
		}

		c.detectedAt = clock.Time{}
		c.id = 0
	}

	if c.Slave != nil {
		c.Slave.Inhibit()
	}
}
