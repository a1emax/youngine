package ebiteninput

import (
	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// Keyboard state based on Ebitengine.
type Keyboard interface {
	input.Keyboard

	// Update updates state.
	Update()
}

// keyboardImpl is the implementation of the [Keyboard] interface.
type keyboardImpl struct {
	keys [input.KeyboardKeyCount]keyboardKeyImpl
}

// NewKeyboard initializes and returns new [Keyboard].
func NewKeyboard(clk clock.Clock) Keyboard {
	if clk == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	k := &keyboardImpl{}
	k.init(clk)

	return k
}

// init initializes [Keyboard].
func (k *keyboardImpl) init(clk clock.Clock) {
	for i := range k.keys {
		k.keys[i].init(input.KeyboardKeyCode(i+1), clk)
	}
}

// Key implements the [input.Keyboard] interface.
func (k *keyboardImpl) Key(code input.KeyboardKeyCode) input.KeyboardKey {
	if code < input.MinKeyboardKeyCode || code > input.MaxKeyboardKeyCode {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	return &k.keys[code-1]
}

// Mark implements the [input.Keyboard] interface.
func (k *keyboardImpl) Mark() {
	for i := range k.keys {
		k.keys[i].Mark()
	}
}

// Update implements the [Keyboard] interface.
func (k *keyboardImpl) Update() {
	for i := range k.keys {
		k.keys[i].Update()
	}
}
