package input

import (
	"github.com/a1emax/youngine/fault"
)

// Sensor handles input with background of type B.
type Sensor[B any] interface {

	// Actuate actuates sensor.
	Actuate(background B)

	// Inhibit inhibits sensor.
	Inhibit()
}

// MultiSensor is [Sensor] composed of multiple other ones.
type MultiSensor[B any] []Sensor[B]

// Actuate implements the [Sensor] interface.
func (s MultiSensor[B]) Actuate(background B) {
	for _, v := range s {
		if v == nil {
			panic(fault.Trace(fault.ErrNilPointer))
		}

		v.Actuate(background)
	}
}

// Inhibit implements the [Sensor] interface.
func (s MultiSensor[B]) Inhibit() {
	for _, v := range s {
		if v == nil {
			panic(fault.Trace(fault.ErrNilPointer))
		}

		v.Inhibit()
	}
}
