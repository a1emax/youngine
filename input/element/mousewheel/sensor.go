package mousewheel

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// Sensor handles wheel input with background of type B.
type Sensor[B any] interface {
	input.Sensor[B]
}

// SensorConfig configures [Sensor].
type SensorConfig[B any] struct {

	// Wheel state.
	Wheel input.MouseWheel

	// OnScroll, if specified, is called on [ScrollEvent].
	OnScroll func(event ScrollEvent[B])

	// Slave, if specified, is actuated if wheel is scrolled, or inhibited otherwise.
	Slave input.Sensor[Background[B]]
}

// sensorImpl is the implementation of the [Sensor] interface.
type sensorImpl[B any] struct {
	SensorConfig[B]
}

// NewSensor initializes and returns new [Sensor].
func NewSensor[B any](config SensorConfig[B]) Sensor[B] {
	if config.Wheel == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &sensorImpl[B]{
		SensorConfig: config,
	}
}

// Actuate implements the [input.Sensor] interface.
func (s *sensorImpl[B]) Actuate(background B) {
	if s.Wheel.IsMarked() {
		s.Inhibit()

		return
	}

	s.Wheel.Mark()

	wheelOffset := s.Wheel.Offset()
	if wheelOffset.IsZero() {
		s.Inhibit()

		return
	}

	if s.OnScroll != nil {
		s.OnScroll(ScrollEvent[B]{
			Background: background,
			Offset:     wheelOffset,
		})
	}

	if s.Slave != nil {
		s.Slave.Actuate(Background[B]{
			Background: background,
			Offset:     wheelOffset,
		})
	}
}

// Inhibit implements the [input.Sensor] interface.
func (s *sensorImpl[B]) Inhibit() {
	if s.Slave != nil {
		s.Slave.Inhibit()
	}
}
