package ebiteninput

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/tempo"
)

// MouseButton state based on Ebitengine.
type MouseButton interface {
	input.MouseButton

	// Update updates state.
	Update()
}

// mouseButtonImpl is the implementation of the [MouseButton] interface.
type mouseButtonImpl struct {
	code       input.MouseButtonCode
	nower      tempo.Nower
	pressedAt  tempo.Time
	releasedAt tempo.Time
	isMarked   bool
}

// NewMouseButton initializes and returns new [MouseButton] with given code.
func NewMouseButton(code input.MouseButtonCode, nower tempo.Nower) MouseButton {
	if nower == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &mouseButtonImpl{
		code:  code,
		nower: nower,
	}
}

// Code implements the [input.MouseButton] interface.
func (b *mouseButtonImpl) Code() input.MouseButtonCode {
	return b.code
}

// PressedAt implements the [input.MouseButton] interface.
func (b *mouseButtonImpl) PressedAt() tempo.Time {
	return b.pressedAt
}

// ReleasedAt implements the [input.MouseButton] interface.
func (b *mouseButtonImpl) ReleasedAt() tempo.Time {
	return b.releasedAt
}

// IsMarked implements the [input.MouseButton] interface.
func (b *mouseButtonImpl) IsMarked() bool {
	return b.isMarked
}

// Mark implements the [input.MouseButton] interface.
func (b *mouseButtonImpl) Mark() {
	b.isMarked = true
}

// Update implements the [MouseButton] interface.
func (b *mouseButtonImpl) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton(b.code - 1)) {
		if b.pressedAt.IsZero() {
			b.pressedAt = b.nower.Now()
		}
	} else {
		if !b.pressedAt.IsZero() {
			b.pressedAt = tempo.Time{}
			b.releasedAt = b.nower.Now()
		}
	}

	b.isMarked = false
}
