//go:build !android && !ios

package ebiteninput

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/basic"
)

// mouseCursorImpl is the implementation of the [MouseCursor] interface.
type mouseCursorImpl struct {
	position basic.Vec2
	isMarked bool
}

// NewMouseCursor initializes and returns new [MouseCursor].
func NewMouseCursor() MouseCursor {
	c := &mouseCursorImpl{}
	c.init()

	return c
}

// init initializes [MouseCursor].
func (c *mouseCursorImpl) init() {
}

// IsAvailable implements the [input.MouseCursor] interface.
func (c *mouseCursorImpl) IsAvailable() bool {
	return true
}

// Position implements the [input.MouseCursor] interface.
func (c *mouseCursorImpl) Position() basic.Vec2 {
	return c.position
}

// IsMarked implements the [input.MouseCursor] interface.
func (c *mouseCursorImpl) IsMarked() bool {
	return c.isMarked
}

// Mark implements the [input.MouseCursor] interface.
func (c *mouseCursorImpl) Mark() {
	c.isMarked = true
}

// Update implements the [MouseCursor] interface.
func (c *mouseCursorImpl) Update() {
	x, y := ebiten.CursorPosition()
	c.position = basic.Vec2{basic.Float(x), basic.Float(y)}

	c.isMarked = false
}
