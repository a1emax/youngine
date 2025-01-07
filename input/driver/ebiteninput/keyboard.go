package ebiteninput

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/tempo"
)

// Keyboard state based on Ebitengine.
type Keyboard interface {
	input.Keyboard

	// Update updates state.
	Update()
}

// keyboardImpl is the implementation of the [Keyboard] interface.
type keyboardImpl struct {
	keys [input.KeyboardKeyCount]KeyboardKey
}

// NewKeyboard initializes and returns new [Keyboard].
func NewKeyboard(nower tempo.Nower) Keyboard {
	if nower == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	k := &keyboardImpl{}

	for i := range k.keys {
		k.keys[i] = NewKeyboardKey(input.KeyboardKeyCode(i+1), nower)
	}

	return k
}

// Key implements the [input.Keyboard] interface.
func (k *keyboardImpl) Key(code input.KeyboardKeyCode) input.KeyboardKey {
	if code < input.MinKeyboardKeyCode || code > input.MaxKeyboardKeyCode {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	return k.keys[code-1]
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
