package padding

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// Padding of element displayed on screen of type S and extended by trait of type T.
type Padding[S, T any] interface {
	scene.Element[S, T]
}

// paddingImpl is the implementation of the [Padding] interface.
type paddingImpl[S, T any] struct {
	scene.BaseElement[S, T, basic.None]

	attrs           scene.Attrs
	containerLayout containerLayout
	element         Element[S]
	elementLayout   elementLayout
}

// New initializes and returns new [Padding].
func New[S, T any](traitFunc scene.TraitFunc[T], element Element[S]) Padding[S, T] {
	if traitFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if element == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	p := &paddingImpl[S, T]{}
	p.Init(traitFunc, func(basic.None) basic.None {
		return basic.None{}
	})

	p.element = element

	return p
}

// IsOff implements the [scene.Element] interface.
func (p *paddingImpl[S, T]) IsOff() bool {
	return p.element.IsOff()
}

// Attrs implements the [scene.Element] interface.
func (p *paddingImpl[S, T]) Attrs() scene.Attrs {
	return p.attrs
}

// Refresh implements the [scene.Element] interface.
func (p *paddingImpl[S, T]) Refresh() {
	p.BaseElement.Refresh()
	p.attrs = scene.Attrs{}

	p.element.Refresh()
}

// Exclude implements the [scene.Element] interface.
func (p *paddingImpl[S, T]) Exclude() {
	p.element.Exclude()
}

// Actuate implements the [scene.Element] interface.
func (p *paddingImpl[S, T]) Actuate() {
	p.element.Actuate()
}

// Inhibit implements the [scene.Element] interface.
func (p *paddingImpl[S, T]) Inhibit() {
	p.element.Inhibit()
}

// Update implements the [scene.Element] interface.
func (p *paddingImpl[S, T]) Update() {
	p.element.Update()
}

// Draw implements the [scene.Element] interface.
func (p *paddingImpl[S, T]) Draw(screen S) {
	p.element.Draw(screen)
}

// Dispose implements the [scene.Element] interface.
func (p *paddingImpl[S, T]) Dispose() {
	p.element.Dispose()
}
