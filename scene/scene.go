package scene

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// Update updates scene with given root element.
func Update[S, T any](root Element[S, T], bbox basic.Rect) {
	if root == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	root.Refresh()

	if root.IsOff() {
		root.Exclude()
		root.Inhibit()

		return
	}

	root.Prepare()

	attrs := root.Attrs()
	minWidth := attrs.MinWidth.Or(0)
	maxWidth := max(minWidth, attrs.MaxWidth.Or(basic.PosInf()))
	minHeight := attrs.MinHeight.Or(0)
	maxHeight := max(minHeight, attrs.MaxHeight.Or(basic.PosInf()))
	root.Arrange(basic.Rect{
		Min: bbox.Min,
		Size: basic.Vec2{
			basic.Clamp(bbox.Width(), minWidth, maxWidth),
			basic.Clamp(bbox.Height(), minHeight, maxHeight),
		},
	})

	root.Actuate()
	root.Update()
}

// Draw draws scene with given root element on given screen.
func Draw[S, T any](root Element[S, T], screen S) {
	if root == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	if root.IsOff() {
		return
	}

	root.Draw(screen)
}

// Dispose disposes scene with given root element.
func Dispose[S, T any](root Element[S, T]) {
	if root == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	root.Dispose()
}
