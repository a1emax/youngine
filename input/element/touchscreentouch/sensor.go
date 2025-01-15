package touchscreentouch

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/tempo"
)

// Sensor handles touch input with background of type B.
type Sensor[B any] interface {
	input.Sensor[B]
}

// SensorConfig configures [Sensor].
type SensorConfig[B any] struct {

	// Touchscreen state.
	Touchscreen input.Touchscreen

	// Nower representing time.
	Nower tempo.Nower

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

	// Slave, if specified, is actuated if touch is inside region, or inhibited otherwise.
	Slave input.Sensor[Background[B]]
}

// sensorImpl is the implementation of the [Sensor] interface.
type sensorImpl[B any] struct {
	SensorConfig[B]

	touchDetectedAt tempo.Time
	touchID         input.TouchscreenTouchID // Valid only if touchDetectedAt is not zero.
}

// NewSensor initializes and returns new [Sensor].
func NewSensor[B any](config SensorConfig[B]) Sensor[B] {
	if config.Touchscreen == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if config.Nower == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &sensorImpl[B]{
		SensorConfig: config,
	}
}

// Actuate implements the [input.Sensor] interface.
func (s *sensorImpl[B]) Actuate(background B) {
	var prevTouchFound bool
	var touch input.TouchscreenTouch
	var touchID input.TouchscreenTouchID
	var touchPosition basic.Vec2
	for i, n := 0, s.Touchscreen.TouchCount(); i < n; i++ {
		currentTouch := s.Touchscreen.Touch(i)

		currentTouchID := currentTouch.ID()
		if !s.touchDetectedAt.IsZero() && currentTouchID == s.touchID {
			prevTouchFound = true
		}

		if currentTouch.IsMarked() {
			continue
		}

		currentTouchPosition := currentTouch.Position()
		if s.HitTest != nil && !s.HitTest(currentTouchPosition) {
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
			if s.OnGone != nil {
				s.OnGone(GoneEvent{})
			}
		} else {
			if s.OnEnd != nil {
				s.OnEnd(EndEvent[B]{
					Background: background,
				})
			}
		}

		s.touchDetectedAt = tempo.Time{}
		s.touchID = 0
	}

	if touch == nil {
		if !s.touchDetectedAt.IsZero() {
			finish()
		}

		if s.Slave != nil {
			s.Slave.Inhibit()
		}

		return
	}

	touch.Mark()

	// Check for different touch.
	if !s.touchDetectedAt.IsZero() && touchID != s.touchID {
		finish()
	}

	now := s.Nower.Now()

	var touchDuration tempo.Ticks
	if s.touchDetectedAt.IsZero() {
		s.touchDetectedAt = now
		s.touchID = touchID

		if touch.StartedAt() == now {
			if s.OnStart != nil {
				s.OnStart(StartEvent[B]{
					Background: background,
				})
			}
		}
	} else {
		touchDuration = now.Sub(s.touchDetectedAt)
	}
	touchDuration++ // Starts from 1.

	if s.OnHover != nil {
		s.OnHover(HoverEvent[B]{
			Background: background,
			Duration:   touchDuration,
			Position:   touchPosition,
		})
	}

	if s.Slave != nil {
		s.Slave.Actuate(Background[B]{
			Background: background,
			Duration:   touchDuration,
			Position:   touchPosition,
		})
	}
}

// Inhibit implements the [input.Sensor] interface.
func (s *sensorImpl[B]) Inhibit() {
	if !s.touchDetectedAt.IsZero() {
		if s.OnGone != nil {
			s.OnGone(GoneEvent{})
		}

		s.touchDetectedAt = tempo.Time{}
		s.touchID = 0
	}

	if s.Slave != nil {
		s.Slave.Inhibit()
	}
}
