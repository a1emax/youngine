package ebiteninput

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
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
	clock      clock.Clock
	pressedAt  clock.Time
	releasedAt clock.Time
	isMarked   bool
}

// NewMouseButton initializes and returns new [MouseButton] with given code.
func NewMouseButton(code input.MouseButtonCode, clk clock.Clock) MouseButton {
	if code < input.MinMouseButtonCode || code > input.MaxMouseButtonCode {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}
	if clk == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	b := &mouseButtonImpl{}
	b.init(code, clk)

	return b
}

// init initializes [MouseButton].
func (b *mouseButtonImpl) init(code input.MouseButtonCode, clk clock.Clock) {
	b.code = code
	b.clock = clk
}

// Code implements the [input.MouseButton] interface.
func (b *mouseButtonImpl) Code() input.MouseButtonCode {
	return b.code
}

// PressedAt implements the [input.MouseButton] interface.
func (b *mouseButtonImpl) PressedAt() clock.Time {
	return b.pressedAt
}

// ReleasedAt implements the [input.MouseButton] interface.
func (b *mouseButtonImpl) ReleasedAt() clock.Time {
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
			b.pressedAt = b.clock.Now()
		}
	} else {
		if !b.pressedAt.IsZero() {
			b.pressedAt = clock.Time{}
			b.releasedAt = b.clock.Now()
		}
	}

	b.isMarked = false
}
