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

	// Touchscreen state.
	Touchscreen input.Touchscreen

	// Clock representing time.
	Clock clock.Clock

	// HitTest, if specified, restricts active area touch is being handled inside.
	HitTest func(position basic.Vec2) bool

	// OnStart, if specified, is called on [StartEvent].
	OnStart func(event StartEvent[B])

	// OnHover, if specified, is called on [HoverEvent].
	OnHover func(event HoverEvent[B])

	// OnEnd, if specified, is called on [EndEvent].
	OnEnd func(event EndEvent[B])

	// OnGone, if specified, is called on [GoneEvent].
	OnGone func(event GoneEvent)

	// Slave, if specified, is actuated if touch is inside active area, or inhibited otherwise.
	Slave input.Controller[Background[B]]
}

// controllerImpl is the implementation of the [Controller] interface.
type controllerImpl[B any] struct {
	ControllerConfig[B]

	touchDetectedAt clock.Time
	touchID         input.TouchscreenTouchID // Valid only if touchDetectedAt is not zero.
}

// NewController initializes and returns new [Controller].
func NewController[B any](config ControllerConfig[B]) Controller[B] {
	if config.Touchscreen == nil {
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
	var prevTouchFound bool
	var touch input.TouchscreenTouch
	var touchID input.TouchscreenTouchID
	var touchPosition basic.Vec2
	for i, n := 0, c.Touchscreen.TouchCount(); i < n; i++ {
		currentTouch := c.Touchscreen.Touch(i)

		currentTouchID := currentTouch.ID()
		if !c.touchDetectedAt.IsZero() && currentTouchID == c.touchID {
			prevTouchFound = true
		}

		if currentTouch.IsMarked() {
			continue
		}

		currentTouchPosition := currentTouch.Position()
		if c.HitTest != nil && !c.HitTest(currentTouchPosition) {
			continue
		}

		// Check for multiple touches.
		if touch != nil {
			touch = nil

			break
		}

		touch = currentTouch
		touchID = currentTouchID
		touchPosition = currentTouchPosition
	}

	finish := func() {
		if prevTouchFound {
			if c.OnGone != nil {
				c.OnGone(GoneEvent{})
			}
		} else {
			if c.OnEnd != nil {
				c.OnEnd(EndEvent[B]{
					Background: background,
				})
			}
		}

		c.touchDetectedAt = clock.Time{}
		c.touchID = 0
	}

	if touch == nil {
		if !c.touchDetectedAt.IsZero() {
			finish()
		}

		if c.Slave != nil {
			c.Slave.Inhibit()
		}

		return
	}

	touch.Mark()

	// Check for different touch.
	if !c.touchDetectedAt.IsZero() && touchID != c.touchID {
		finish()
	}

	now := c.Clock.Now()

	var touchDuration clock.Ticks
	if c.touchDetectedAt.IsZero() {
		c.touchDetectedAt = now
		c.touchID = touchID

		if touch.StartedAt() == now {
			if c.OnStart != nil {
				c.OnStart(StartEvent[B]{
					Background: background,
				})
			}
		}
	} else {
		touchDuration = now.Sub(c.touchDetectedAt)
	}
	touchDuration++ // Starts from 1.

	if c.OnHover != nil {
		c.OnHover(HoverEvent[B]{
			Background: background,
			Duration:   touchDuration,
			Position:   touchPosition,
		})
	}

	if c.Slave != nil {
		c.Slave.Actuate(Background[B]{
			Background: background,
			Duration:   touchDuration,
			Position:   touchPosition,
		})
	}
}

// Inhibit implements the [input.Controller] interface.
func (c *controllerImpl[B]) Inhibit() {
	if !c.touchDetectedAt.IsZero() {
		if c.OnGone != nil {
			c.OnGone(GoneEvent{})
		}

		c.touchDetectedAt = clock.Time{}
		c.touchID = 0
	}

	if c.Slave != nil {
		c.Slave.Inhibit()
	}
}
