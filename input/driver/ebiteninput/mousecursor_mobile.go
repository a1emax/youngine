//go:build android || ios

package ebiteninput

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// mouseCursorImpl is the implementation of the [MouseCursor] interface.
type mouseCursorImpl struct {
	isMarked bool
}

// IsAvailable implements the [input.MouseCursor] interface.
func (c *mouseCursorImpl) IsAvailable() bool {
	return false
}

// Position implements the [input.MouseCursor] interface.
func (c *mouseCursorImpl) Position() basic.Vec2 {
	panic(fault.Trace(fault.ErrInvalidUse))
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
	c.isMarked = false
}
