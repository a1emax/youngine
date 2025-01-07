package ebiteninput

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/tempo"
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
	nower      tempo.Nower
	pressedAt  tempo.Time
	releasedAt tempo.Time
	isMarked   bool
}

// NewKeyboardKey initializes and returns new [KeyboardKey] with given code.
func NewKeyboardKey(code input.KeyboardKeyCode, nower tempo.Nower) KeyboardKey {
	if nower == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &keyboardKeyImpl{
		code:  code,
		nower: nower,
	}
}

// Code implements the [input.KeyboardKey] interface.
func (k *keyboardKeyImpl) Code() input.KeyboardKeyCode {
	return k.code
}

// PressedAt implements the [input.KeyboardKey] interface.
func (k *keyboardKeyImpl) PressedAt() tempo.Time {
	return k.pressedAt
}

// ReleasedAt implements the [input.KeyboardKey] interface.
func (k *keyboardKeyImpl) ReleasedAt() tempo.Time {
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
			k.pressedAt = k.nower.Now()
		}
	} else {
		if !k.pressedAt.IsZero() {
			k.pressedAt = tempo.Time{}
			k.releasedAt = k.nower.Now()
		}
	}

	k.isMarked = false
}
