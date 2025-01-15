package textbox

import (
	"golang.org/x/image/font"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// StringDrawer draws strings.
type StringDrawer interface {

	// DrawString draws given string starting from (x, y) using given font face.
	DrawString(s string, x, y basic.Float, fontFace font.Face)
}

// StringDrawerFunc is the functional implementation of the [StringDrawer] interface.
type StringDrawerFunc func(s string, x, y basic.Float, fontFace font.Face)

// DrawString implements the [StringDrawer] interface.
func (f StringDrawerFunc) DrawString(s string, x, y basic.Float, fontFace font.Face) {
	if f == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	f(s, x, y, fontFace)
}
