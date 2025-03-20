package ebiteninput

import (
	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// Input state based on Ebitengine.
type Input interface {
	input.Input

	// Update updates state.
	Update()
}

// inputImpl is the implementation of the [Input] interface.
type inputImpl struct {
	keyboard    keyboardImpl
	mouse       mouseImpl
	touchscreen touchscreenImpl
}

// New initializes and returns new [Input].
func New(clk clock.Clock) Input {
	if clk == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	i := &inputImpl{}
	i.keyboard.init(clk)
	i.mouse.init(clk)
	i.touchscreen.init(clk)

	return i
}

// Keyboard implements the [input.Input] interface.
func (i *inputImpl) Keyboard() input.Keyboard {
	return &i.keyboard
}

// Mouse implements the [input.Input] interface.
func (i *inputImpl) Mouse() input.Mouse {
	return &i.mouse
}

// Touchscreen implements the [input.Input] interface.
func (i *inputImpl) Touchscreen() input.Touchscreen {
	return &i.touchscreen
}

// Mark implements the [input.Input] interface.
func (i *inputImpl) Mark() {
	i.keyboard.Mark()
	i.mouse.Mark()
	i.touchscreen.Mark()
}

// Update implements the [Input] interface.
func (i *inputImpl) Update() {
	i.keyboard.Update()
	i.mouse.Update()
	i.touchscreen.Update()
}
