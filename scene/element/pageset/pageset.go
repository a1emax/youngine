package pageset

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// PageSet displayed on screen ot type S and extended by trait or type T.
type PageSet[S, T any] interface {
	scene.Element[S, T]
}

// Props associated with [PageSet].
type Props struct {

	// Page specifies index of current page.
	Page int
}

// Func returns these props.
func (p Props) Func(Props) Props {
	return p
}

// pageSetImpl is the implementation of the [PageSet] interface.
type pageSetImpl[S, T any] struct {
	scene.BaseElement[S, T, Props]

	pages []Page[S]
}

// New initializes and returns new [PageSet].
func New[S, T any](traitFunc scene.TraitFunc[T], propsFunc scene.PropsFunc[Props], pages ...Page[S]) PageSet[S, T] {
	if traitFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if propsFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	p := &pageSetImpl[S, T]{}
	p.Init(traitFunc, propsFunc)

	p.pages = make([]Page[S], 0, len(pages))
	for _, page := range pages {
		if page == nil {
			panic(fault.Trace(fault.ErrNilPointer))
		}

		p.pages = append(p.pages, page)
	}

	return p
}

// IsOff implements the [scene.Element] interface.
func (p *pageSetImpl[S, T]) IsOff() bool {
	pi := p.Props().Page
	if pi < 0 || pi >= len(p.pages) {
		return true
	}

	return p.pages[pi].IsOff()
}

// Attrs implements the [scene.Element] interface.
func (p *pageSetImpl[S, T]) Attrs() scene.Attrs {
	pi := p.Props().Page
	if pi < 0 || pi >= len(p.pages) {
		return scene.Attrs{}
	}

	return p.pages[pi].Attrs()
}

// Refresh implements the [scene.Element] interface.
func (p *pageSetImpl[S, T]) Refresh() {
	p.BaseElement.Refresh()

	for _, page := range p.pages {
		page.Refresh()
	}
}

// Prepare implements the [scene.Element] interface.
func (p *pageSetImpl[S, T]) Prepare() {
	pi := p.Props().Page
	for i, page := range p.pages {
		if i == pi {
			page.Prepare()
		} else {
			page.Exclude()
		}
	}
}

// Exclude implements the [scene.Element] interface.
func (p *pageSetImpl[S, T]) Exclude() {
	for _, page := range p.pages {
		page.Exclude()
	}
}

// Arrange implements the [scene.Element] interface.
func (p *pageSetImpl[S, T]) Arrange(bbox basic.Rect) {
	pi := p.Props().Page
	if pi < 0 || pi >= len(p.pages) {
		return
	}

	p.pages[pi].Arrange(bbox)
}

// Actuate implements the [scene.Element] interface.
func (p *pageSetImpl[S, T]) Actuate() {
	pi := p.Props().Page
	for i, page := range p.pages {
		if i == pi {
			page.Actuate()
		} else {
			page.Inhibit()
		}
	}
}

// Inhibit implements the [scene.Element] interface.
func (p *pageSetImpl[S, T]) Inhibit() {
	for _, page := range p.pages {
		page.Inhibit()
	}
}

// Update implements the [scene.Element] interface.
func (p *pageSetImpl[S, T]) Update() {
	pi := p.Props().Page
	if pi < 0 || pi >= len(p.pages) {
		return
	}

	p.pages[pi].Update()
}

// Draw implements the [scene.Element] interface.
func (p *pageSetImpl[S, T]) Draw(screen S) {
	pi := p.Props().Page
	if pi < 0 || pi >= len(p.pages) {
		return
	}

	p.pages[pi].Draw(screen)
}

// Dispose implements the [scene.Element] interface.
func (p *pageSetImpl[S, T]) Dispose() {
	for _, page := range p.pages {
		page.Dispose()
	}
}
