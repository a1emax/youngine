package pageset

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// PageSet placed on screen ot type S inside region or type R.
type PageSet[S any, R scene.Region] interface {
	scene.Element[S, R]
}

// Config configures [PageSet].
type Config struct {

	// StateFunc accepts current state and returns new one.
	StateFunc func(state State) State
}

// State is changeable state of [Overlay].
type State struct {

	// Page specifies index of current page.
	Page int
}

// pageSetImpl is the implementation of the [PageSet] interface.
type pageSetImpl[S any, R scene.Region] struct {
	scene.BaseElement[S, R]
	Config

	region R
	state  State
	pages  []Page[S]
}

// New initializes and returns new [PageSet].
func New[S any, R scene.Region](region R, config Config, pages ...Page[S]) PageSet[S, R] {
	if config.StateFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	pagesCopy := make([]Page[S], 0, len(pages))
	for _, page := range pages {
		if page == nil {
			panic(fault.Trace(fault.ErrNilPointer))
		}

		pagesCopy = append(pagesCopy, page)
	}

	return &pageSetImpl[S, R]{
		Config: config,

		region: region,
		pages:  pagesCopy,
	}
}

// Region implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) Region() R {
	return p.region
}

// IsActive implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) IsActive() bool {
	if p.state.Page < 0 || p.state.Page >= len(p.pages) {
		return false
	}

	return p.pages[p.state.Page].IsActive()
}

// Outline implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) Outline() scene.Outline {
	if p.state.Page < 0 || p.state.Page >= len(p.pages) {
		return scene.Outline{}
	}

	return p.pages[p.state.Page].Outline()
}

// Refresh implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) Refresh() {
	p.state = p.StateFunc(p.state)

	for _, page := range p.pages {
		page.Refresh()
	}
}

// Prepare implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) Prepare() {
	for i, page := range p.pages {
		if i == p.state.Page {
			page.Prepare()
		} else {
			page.Exclude()
		}
	}
}

// Exclude implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) Exclude() {
	for _, page := range p.pages {
		page.Exclude()
	}
}

// Arrange implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) Arrange() {
	if p.state.Page < 0 || p.state.Page >= len(p.pages) {
		return
	}

	page := p.pages[p.state.Page]
	page.Region().Arrange(p.region.Rect())
	page.Arrange()
}

// Actuate implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) Actuate() {
	for i, page := range p.pages {
		if i == p.state.Page {
			page.Actuate()
		} else {
			page.Inhibit()
		}
	}
}

// Inhibit implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) Inhibit() {
	for _, page := range p.pages {
		page.Inhibit()
	}
}

// Update implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) Update() {
	if p.state.Page < 0 || p.state.Page >= len(p.pages) {
		return
	}

	p.pages[p.state.Page].Update()
}

// Draw implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) Draw(screen S) {
	if p.state.Page < 0 || p.state.Page >= len(p.pages) {
		return
	}

	p.pages[p.state.Page].Draw(screen)
}

// Dispose implements the [scene.Element] interface.
func (p *pageSetImpl[S, R]) Dispose() {
	for _, page := range p.pages {
		page.Dispose()
	}
}
