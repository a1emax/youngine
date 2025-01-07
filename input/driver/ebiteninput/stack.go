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

// stackImpl is the implementation of the [Stack] interface.
type stackImpl struct {
	keyboard    Keyboard
	mouse       Mouse
	touchscreen Touchscreen
}

// NewStack initializes and returns new [Stack].
func NewStack(nower tempo.Nower) Stack {
	if nower == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &stackImpl{
		keyboard:    NewKeyboard(nower),
		mouse:       NewMouse(nower),
		touchscreen: NewTouchscreen(nower),
	}
}

// Keyboard implements the [input.Stack] interface.
func (s *stackImpl) Keyboard() input.Keyboard {
	return s.keyboard
}

// Mouse implements the [input.Stack] interface.
func (s *stackImpl) Mouse() input.Mouse {
	return s.mouse
}

// Touchscreen implements the [input.Stack] interface.
func (s *stackImpl) Touchscreen() input.Touchscreen {
	return s.touchscreen
}

// Mark implements the [input.Stack] interface.
func (s *stackImpl) Mark() {
	s.keyboard.Mark()
	s.mouse.Mark()
	s.touchscreen.Mark()
}

// Update implements the [Stack] interface.
func (s *stackImpl) Update() {
	s.keyboard.Update()
	s.mouse.Update()
	s.touchscreen.Update()
}
