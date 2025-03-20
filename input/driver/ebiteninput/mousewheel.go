package ebiteninput

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input"
)

// MouseWheel state based on Ebitengine.
type MouseWheel interface {
	input.MouseWheel

	// Update updates state.
	Update()
}

// mouseWheelImpl is the implementation of the [MouseWheel] interface.
type mouseWheelImpl struct {
	offset   basic.Vec2
	isMarked bool
}

// NewMouseWheel initializes and returns new [MouseWheel].
func NewMouseWheel() MouseWheel {
	w := &mouseWheelImpl{}
	w.init()

	return w
}

// init initializes [MouseWheel].
func (w *mouseWheelImpl) init() {
}

// Offset implements the [input.MouseWheel] interface.
func (w *mouseWheelImpl) Offset() basic.Vec2 {
	return w.offset
}

// IsMarked implements the [input.MouseWheel] interface.
func (w *mouseWheelImpl) IsMarked() bool {
	return w.isMarked
}

// Mark implements the [input.MouseWheel] interface.
func (w *mouseWheelImpl) Mark() {
	w.isMarked = true
}

// Update implements the [MouseWheel] interface.
func (w *mouseWheelImpl) Update() {
	xoff, yoff := ebiten.Wheel()
	w.offset = basic.Vec2{xoff, yoff}

	w.isMarked = false
}
