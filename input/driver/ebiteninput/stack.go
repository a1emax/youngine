package ebiteninput

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/tempo"
)

// Stack state based on Ebitengine.
type Stack interface {
	input.Stack

	// Update updates state.
	Update()
}

// systemImpl is the implementation of the [Stack] interface.
type systemImpl struct {
	keyboard    Keyboard
	mouse       Mouse
	touchscreen Touchscreen
}

// NewSystem initializes and returns new [Stack].
func NewSystem(nower tempo.Nower) Stack {
	if nower == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &systemImpl{
		keyboard:    NewKeyboard(nower),
		mouse:       NewMouse(nower),
		touchscreen: NewTouchscreen(nower),
	}
}

// Keyboard implements the [input.Stack] interface.
func (s *systemImpl) Keyboard() input.Keyboard {
	return s.keyboard
}

// Mouse implements the [input.Stack] interface.
func (s *systemImpl) Mouse() input.Mouse {
	return s.mouse
}

// Touchscreen implements the [input.Stack] interface.
func (s *systemImpl) Touchscreen() input.Touchscreen {
	return s.touchscreen
}

// Mark implements the [input.Stack] interface.
func (s *systemImpl) Mark() {
	s.keyboard.Mark()
	s.mouse.Mark()
	s.touchscreen.Mark()
}

// Update implements the [Stack] interface.
func (s *systemImpl) Update() {
	s.keyboard.Update()
	s.mouse.Update()
	s.touchscreen.Update()
}
