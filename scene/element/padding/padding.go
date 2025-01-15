package padding

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// Padding of element placed on screen of type S inside region of type R.
type Padding[S any, R scene.Region] interface {
	scene.Element[S, R]
}

// paddingImpl is the implementation of the [Padding] interface.
type paddingImpl[S any, R scene.Region] struct {
	scene.BaseElement[S, R]

	region  R
	element Element[S]

	outline         scene.Outline
	containerLayout containerLayout
	elementLayout   elementLayout
}

// New initializes and returns new [Padding].
func New[S any, R scene.Region](region R, element Element[S]) Padding[S, R] {
	if element == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &paddingImpl[S, R]{
		region:  region,
		element: element,
	}
}

// Region implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) Region() R {
	return p.region
}

// IsActive implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) IsActive() bool {
	return p.element.IsActive()
}

// Outline implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) Outline() scene.Outline {
	return p.outline
}

// Refresh implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) Refresh() {
	p.outline = scene.Outline{}

	p.element.Region().Refresh()
	p.element.Refresh()
}

// Exclude implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) Exclude() {
	p.element.Exclude()
}

// Actuate implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) Actuate() {
	p.element.Actuate()
}

// Inhibit implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) Inhibit() {
	p.element.Inhibit()
}

// Update implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) Update() {
	p.element.Update()
}

// Draw implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) Draw(screen S) {
	p.element.Draw(screen)
}

// Dispose implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) Dispose() {
	p.element.Dispose()
}
