package mousecursor

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// Sensor handles cursor input with background of type B.
type Sensor[B any] interface {
	input.Sensor[B]
}

// SensorConfig configures [Sensor].
type SensorConfig[B any] struct {

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

	// Slave, if specified, is actuated if cursor is inside region, or inhibited otherwise.
	Slave input.Sensor[Background[B]]
}

// sensorImpl is the implementation of the [Sensor] interface.
type sensorImpl[B any] struct {
	SensorConfig[B]

	cursorDetected bool
}

// NewSensor initializes and returns new [Sensor].
func NewSensor[B any](config SensorConfig[B]) Sensor[B] {
	if config.Cursor == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &sensorImpl[B]{
		SensorConfig: config,
	}
}

// Actuate implements the [input.Sensor] interface.
func (s *sensorImpl[B]) Actuate(background B) {
	if !s.Cursor.IsAvailable() || s.Cursor.IsMarked() {
		s.Inhibit()

		return
	}

	cursorPosition := s.Cursor.Position()
	if s.HitTest != nil && !s.HitTest(cursorPosition) {
		s.Inhibit()

		return
	}

	s.Cursor.Mark()

	if !s.cursorDetected {
		s.cursorDetected = true

		if s.OnEnter != nil {
			s.OnEnter(EnterEvent[B]{
				Background: background,
				Position:   cursorPosition,
			})
		}
	}

	if s.OnHover != nil {
		s.OnHover(HoverEvent[B]{
			Background: background,
			Position:   cursorPosition,
		})
	}

	if s.Slave != nil {
		s.Slave.Actuate(Background[B]{
			Background: background,
			Position:   cursorPosition,
		})
	}
}

// Inhibit implements the [input.Sensor] interface.
func (s *sensorImpl[B]) Inhibit() {
	if s.cursorDetected {
		if s.OnLeave != nil {
			s.OnLeave(LeaveEvent{})
		}

		s.cursorDetected = false
	}

	if s.Slave != nil {
		s.Slave.Inhibit()
	}
}
