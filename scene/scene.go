package scene

import (
	"github.com/a1emax/youngine/fault"
)

// Update updates scene with given root element.
func Update[S any, R Region](root Element[S, R]) {
	if root == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	root.Refresh()

	if !root.IsActive() {
		root.Exclude()
		root.Inhibit()

		return
	}

	root.Prepare()
	root.Arrange()
	root.Actuate()
	root.Update()
}

// Draw draws scene with given root element on given screen.
func Draw[S any, R Region](root Element[S, R], screen S) {
	if root == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	if !root.IsActive() {
		return
	}

	root.Draw(screen)
}

// Dispose disposes scene with given root element.
func Dispose[S any, R Region](root Element[S, R]) {
	if root == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	root.Dispose()
}
