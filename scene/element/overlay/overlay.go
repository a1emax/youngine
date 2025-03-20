package overlay

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// Overlay displayed on screen of type S and extended by trait of type T.
type Overlay[S, T any] interface {
	scene.Element[S, T]
}

// Props associated with [Overlay].
type Props struct {
	scene.Attrs
}

// Func returns these props.
func (p Props) Func(Props) Props {
	return p
}

// overlayImpl is the implementation of the [Overlay] interface.
type overlayImpl[S, T any] struct {
	scene.BaseElement[S, T, Props]

	attrs           scene.Attrs
	containerLayout containerLayout
	items           []Item[S]
	itemLayouts     []itemLayout
}

// New initializes and returns new [Overlay].
func New[S, T any](traitFunc scene.TraitFunc[T], propsFunc scene.PropsFunc[Props], items ...Item[S]) Overlay[S, T] {
	if traitFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if propsFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	o := &overlayImpl[S, T]{}
	o.Init(traitFunc, propsFunc)

	o.items = make([]Item[S], 0, len(items))
	for _, item := range items {
		if item == nil {
			panic(fault.Trace(fault.ErrNilPointer))
		}

		o.items = append(o.items, item)
	}

	o.itemLayouts = make([]itemLayout, 0, len(o.items))

	return o
}

// Attrs implements the [scene.Element] interface.
func (o *overlayImpl[S, T]) Attrs() scene.Attrs {
	return o.attrs
}

// Refresh implements the [scene.Element] interface.
func (o *overlayImpl[S, T]) Refresh() {
	o.BaseElement.Refresh()
	o.attrs = scene.Attrs{}

	for _, item := range o.items {
		item.Refresh()
	}
}

// Exclude implements the [scene.Element] interface.
func (o *overlayImpl[S, T]) Exclude() {
	for _, item := range o.items {
		item.Exclude()
	}
}

// Actuate implements the [scene.Element] interface.
func (o *overlayImpl[S, T]) Actuate() {
	for i := len(o.items) - 1; i >= 0; i-- {
		item := o.items[i]

		if item.IsOff() {
			item.Inhibit()
		} else {
			item.Actuate()
		}
	}
}

// Inhibit implements the [scene.Element] interface.
func (o *overlayImpl[S, T]) Inhibit() {
	for i := len(o.items) - 1; i >= 0; i-- {
		item := o.items[i]

		item.Inhibit()
	}
}

// Update implements the [scene.Element] interface.
func (o *overlayImpl[S, T]) Update() {
	for _, item := range o.items {
		if item.IsOff() {
			continue
		}

		item.Update()
	}
}

// Draw implements the [scene.Element] interface.
func (o *overlayImpl[S, T]) Draw(screen S) {
	for _, item := range o.items {
		if item.IsOff() {
			continue
		}

		item.Draw(screen)
	}
}

// Dispose implements the [scene.Element] interface.
func (o *overlayImpl[S, T]) Dispose() {
	for _, item := range o.items {
		item.Dispose()
	}
}
