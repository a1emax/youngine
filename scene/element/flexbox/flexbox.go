package flexbox

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// Flexbox displayed on screen of type S and extended by trait of type T.
type Flexbox[S, T any] interface {
	scene.Element[S, T]
}

// Props associated with [Flexbox].
type Props struct {
	scene.Attrs

	// AlignItems specifies alignment of items on container's cross axis.
	AlignItems Align

	// Direction specifies how items are placed in container defining main axis.
	Direction Direction

	// JustifyContent specifies how to distribute space between and around items along container's main axis.
	JustifyContent Justify
}

// Func returns these props.
func (p Props) Func(Props) Props {
	return p
}

// flexboxImpl is the implementation of the [Flexbox] interface.
type flexboxImpl[S, T any] struct {
	scene.BaseElement[S, T, Props]

	attrs           scene.Attrs
	containerLayout containerLayout
	items           []Item[S]
	itemLayouts     []itemLayout
}

// New initializes and returns new [Flexbox].
func New[S, T any](traitFunc scene.TraitFunc[T], propsFunc scene.PropsFunc[Props], items ...Item[S]) Flexbox[S, T] {
	if traitFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if propsFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	f := &flexboxImpl[S, T]{}
	f.Init(traitFunc, propsFunc)

	f.items = make([]Item[S], 0, len(items))
	for _, item := range items {
		if item == nil {
			panic(fault.Trace(fault.ErrNilPointer))
		}

		f.items = append(f.items, item)
	}

	f.itemLayouts = make([]itemLayout, 0, len(f.items))

	return f
}

// Attrs implements the [scene.Element] interface.
func (f *flexboxImpl[S, T]) Attrs() scene.Attrs {
	return f.attrs
}

// Refresh implements the [scene.Element] interface.
func (f *flexboxImpl[S, T]) Refresh() {
	f.BaseElement.Refresh()
	f.attrs = scene.Attrs{}

	for _, item := range f.items {
		item.Refresh()
	}
}

// Exclude implements the [scene.Element] interface.
func (f *flexboxImpl[S, T]) Exclude() {
	for _, item := range f.items {
		item.Exclude()
	}
}

// Actuate implements the [scene.Element] interface.
func (f *flexboxImpl[S, T]) Actuate() {
	for i := len(f.items) - 1; i >= 0; i-- {
		item := f.items[i]

		if item.IsOff() {
			item.Inhibit()
		} else {
			item.Actuate()
		}
	}
}

// Inhibit implements the [scene.Element] interface.
func (f *flexboxImpl[S, T]) Inhibit() {
	for i := len(f.items) - 1; i >= 0; i-- {
		item := f.items[i]

		item.Inhibit()
	}
}

// Update implements the [scene.Element] interface.
func (f *flexboxImpl[S, T]) Update() {
	for _, item := range f.items {
		if item.IsOff() {
			continue
		}

		item.Update()
	}
}

// Draw implements the [scene.Element] interface.
func (f *flexboxImpl[S, T]) Draw(screen S) {
	for _, item := range f.items {
		if item.IsOff() {
			continue
		}

		item.Draw(screen)
	}
}

// Dispose implements the [scene.Element] interface.
func (f *flexboxImpl[S, T]) Dispose() {
	for _, item := range f.items {
		item.Dispose()
	}
}
