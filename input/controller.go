package input

import (
	"github.com/a1emax/youngine/fault"
)

// Controller of input element with background of type B.
type Controller[B any] interface {

	// Actuate actuates controller.
	Actuate(background B)

	// Inhibit inhibits controller.
	Inhibit()
}

// MultiController is [Controller] composed of multiple other ones.
type MultiController[B any] []Controller[B]

// Actuate implements the [Controller] interface.
func (c MultiController[B]) Actuate(background B) {
	for _, v := range c {
		if v == nil {
			panic(fault.Trace(fault.ErrNilPointer))
		}

		v.Actuate(background)
	}
}

// Inhibit implements the [Controller] interface.
func (c MultiController[B]) Inhibit() {
	for _, v := range c {
		if v == nil {
			panic(fault.Trace(fault.ErrNilPointer))
		}

		v.Inhibit()
	}
}
