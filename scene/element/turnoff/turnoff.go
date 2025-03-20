package turnoff

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// Turnoff of element displayed on screen of type S and extended by trait of type T.
type Turnoff[S, T any] interface {
	scene.Element[S, T]
}

// Props associated with [Turnoff].
type Props struct {

	// IsOff specifies whether element is off.
	IsOff bool
}

// Func returns these props.
func (p Props) Func(Props) Props {
	return p
}

// turnoffImpl is the implementation of the [Turnoff] interface.
type turnoffImpl[S, T any] struct {
	scene.BaseElement[S, T, Props]

	element Element[S]
}

// New initializes and returns new [Turnoff].
func New[S, T any](traitFunc scene.TraitFunc[T], propsFunc scene.PropsFunc[Props], element Element[S]) Turnoff[S, T] {
	if traitFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if propsFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if element == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	t := &turnoffImpl[S, T]{}
	t.Init(traitFunc, propsFunc)

	t.element = element

	return t
}

// IsOff implements the [scene.Element] interface.
func (t *turnoffImpl[S, T]) IsOff() bool {
	return t.Props().IsOff || t.element.IsOff()
}

// Attrs implements the [scene.Element] interface.
func (t *turnoffImpl[S, T]) Attrs() scene.Attrs {
	return t.element.Attrs()
}

// Refresh implements the [scene.Element] interface.
func (t *turnoffImpl[S, T]) Refresh() {
	t.BaseElement.Refresh()

	t.element.Refresh()
}

// Prepare implements the [scene.Element] interface.
func (t *turnoffImpl[S, T]) Prepare() {
	t.element.Prepare()
}

// Exclude implements the [scene.Element] interface.
func (t *turnoffImpl[S, T]) Exclude() {
	t.element.Exclude()
}

// Arrange implements the [scene.Element] interface.
func (t *turnoffImpl[S, T]) Arrange(bbox basic.Rect) {
	t.element.Arrange(bbox)
}

// Actuate implements the [scene.Element] interface.
func (t *turnoffImpl[S, T]) Actuate() {
	t.element.Actuate()
}

// Inhibit implements the [scene.Element] interface.
func (t *turnoffImpl[S, T]) Inhibit() {
	t.element.Inhibit()
}

// Update implements the [scene.Element] interface.
func (t *turnoffImpl[S, T]) Update() {
	t.element.Update()
}

// Draw implements the [scene.Element] interface.
func (t *turnoffImpl[S, T]) Draw(screen S) {
	t.element.Draw(screen)
}

// Dispose implements the [scene.Element] interface.
func (t *turnoffImpl[S, T]) Dispose() {
	t.element.Dispose()
}
