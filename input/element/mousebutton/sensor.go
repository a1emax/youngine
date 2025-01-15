package mousebutton

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/tempo"
)

// Sensor handles button input with background of type B.
type Sensor[B any] interface {
	input.Sensor[B]
}

// SensorConfig configures [Sensor].
type SensorConfig[B any] struct {

	// Button state.
	Button input.MouseButton

	// Nower representing time.
	Nower tempo.Nower

	// OnDown, if specified, is called on [DownEvent].
	OnDown func(event DownEvent[B])

	// OnPress, if specified, is called on [PressEvent].
	OnPress func(event PressEvent[B])

	// OnUp, if specified, is called on [UpEvent].
	OnUp func(event UpEvent[B])

	// OnGone, if specified, is called on [GoneEvent].
	OnGone func(event GoneEvent)

	// Slave, if specified, is actuated if button is pressed, or inhibited otherwise.
	Slave input.Sensor[Background[B]]
}

// sensorImpl is the implementation of the [Sensor] interface.
type sensorImpl[B any] struct {
	SensorConfig[B]

	buttonPressDetectedAt tempo.Time
}

// NewSensor initializes and returns new [Sensor].
func NewSensor[B any](config SensorConfig[B]) Sensor[B] {
	if config.Button == nil {
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
	if s.Button.IsMarked() {
		s.Inhibit()

		return
	}

	s.Button.Mark()

	now := s.Nower.Now()

	if buttonPressedAt := s.Button.PressedAt(); !buttonPressedAt.IsZero() {
		var buttonPressDuration tempo.Ticks
		if s.buttonPressDetectedAt.IsZero() {
			s.buttonPressDetectedAt = now

			if buttonPressedAt == now {
				if s.OnDown != nil {
					s.OnDown(DownEvent[B]{
						Background: background,
					})
				}
			}
		} else {
			buttonPressDuration = now.Sub(s.buttonPressDetectedAt)
		}
		buttonPressDuration++ // Starts from 1.

		if s.OnPress != nil {
			s.OnPress(PressEvent[B]{
				Background: background,
				Duration:   buttonPressDuration,
			})
		}

		if s.Slave != nil {
			s.Slave.Actuate(Background[B]{
				Background: background,
				Duration:   buttonPressDuration,
			})
		}
	} else {
		if s.Button.ReleasedAt() == now {
			if s.OnUp != nil {
				s.OnUp(UpEvent[B]{
					Background: background,
				})
			}
		}

		s.buttonPressDetectedAt = tempo.Time{}

		if s.Slave != nil {
			s.Slave.Inhibit()
		}
	}
}

// Inhibit implements the [input.Sensor] interface.
func (s *sensorImpl[B]) Inhibit() {
	if !s.buttonPressDetectedAt.IsZero() {
		if s.OnGone != nil {
			s.OnGone(GoneEvent{})
		}
	}

	s.buttonPressDetectedAt = tempo.Time{}

	if s.Slave != nil {
		s.Slave.Inhibit()
	}
}
