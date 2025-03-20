package ebiteninput

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// KeyboardKey state based on Ebitengine.
type KeyboardKey interface {
	input.KeyboardKey

	// Update updates state.
	Update()
}

// keyboardKeyImpl is the implementation of the [KeyboardKey] interface.
type keyboardKeyImpl struct {
	code       input.KeyboardKeyCode
	clock      clock.Clock
	pressedAt  clock.Time
	releasedAt clock.Time
	isMarked   bool
}

// NewKeyboardKey initializes and returns new [KeyboardKey] with given code.
func NewKeyboardKey(code input.KeyboardKeyCode, clk clock.Clock) KeyboardKey {
	if code < input.MinKeyboardKeyCode || code > input.MaxKeyboardKeyCode {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}
	if clk == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	k := &keyboardKeyImpl{}
	k.init(code, clk)

	return k
}

// init initializes [KeyboardKey].
func (k *keyboardKeyImpl) init(code input.KeyboardKeyCode, clk clock.Clock) {
	k.code = code
	k.clock = clk
}

// Code implements the [input.KeyboardKey] interface.
func (k *keyboardKeyImpl) Code() input.KeyboardKeyCode {
	return k.code
}

// PressedAt implements the [input.KeyboardKey] interface.
func (k *keyboardKeyImpl) PressedAt() clock.Time {
	return k.pressedAt
}

// ReleasedAt implements the [input.KeyboardKey] interface.
func (k *keyboardKeyImpl) ReleasedAt() clock.Time {
	return k.releasedAt
}

// IsMarked implements the [input.KeyboardKey] interface.
func (k *keyboardKeyImpl) IsMarked() bool {
	return k.isMarked
}

// Mark implements the [input.KeyboardKey] interface.
func (k *keyboardKeyImpl) Mark() {
	k.isMarked = true
}

// Update implements the [KeyboardKey] interface.
func (k *keyboardKeyImpl) Update() {
	if ebiten.IsKeyPressed(ebiten.Key(k.code - 1)) {
		if k.pressedAt.IsZero() {
			k.pressedAt = k.clock.Now()
		}
	} else {
		if !k.pressedAt.IsZero() {
			k.pressedAt = clock.Time{}
			k.releasedAt = k.clock.Now()
		}
	}

	k.isMarked = false
}
