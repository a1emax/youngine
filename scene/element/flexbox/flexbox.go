package flexbox

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// Flexbox placed on screen of type S inside region of type R.
type Flexbox[S any, R scene.Region] interface {
	scene.Element[S, R]
}

// Config configures [Flexbox].
type Config struct {

	// StateFunc accepts current state and returns new one.
	StateFunc func(state State) State
}

// State is changeable state of [Flexbox].
type State struct {

	// IsInactive specifies whether element is inactive.
	IsInactive bool

	// Outline specifies base outline of element.
	scene.Outline

	// Direction specifies how items are placed in container defining main axis.
	Direction Direction

	// JustifyContent specifies how to distribute space between and around items along container's main axis.
	JustifyContent Justify

	// AlignItems specifies alignment of items on container's cross axis.
	AlignItems Align
}

// flexboxImpl is the implementation of the [Flexbox] interface.
type flexboxImpl[S any, R scene.Region] struct {
	scene.BaseElement[S, R]
	Config

	region R
	state  State
	items  []Item[S]

	outline         scene.Outline
	containerLayout containerLayout
	itemLayouts     []itemLayout
}

// New initializes and returns new [Flexbox].
func New[S any, R scene.Region](region R, config Config, items ...Item[S]) Flexbox[S, R] {
	if config.StateFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	f := &flexboxImpl[S, R]{
		Config: config,

		region: region,
	}

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

// Region implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) Region() R {
	return f.region
}

// IsActive implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) IsActive() bool {
	return !f.state.IsInactive
}

// Outline implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) Outline() scene.Outline {
	return f.outline
}

// Refresh implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) Refresh() {
	f.state = f.StateFunc(f.state)
	f.outline = scene.Outline{}

	for _, item := range f.items {
		item.Region().Refresh()
		item.Refresh()
	}
}

// Exclude implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) Exclude() {
	for _, item := range f.items {
		item.Exclude()
	}
}

// Actuate implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) Actuate() {
	for i := len(f.items) - 1; i >= 0; i-- {
		item := f.items[i]

		if item.IsActive() {
			item.Actuate()
		} else {
			item.Inhibit()
		}
	}
}

// Inhibit implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) Inhibit() {
	for i := len(f.items) - 1; i >= 0; i-- {
		item := f.items[i]

		item.Inhibit()
	}
}

// Update implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) Update() {
	for _, item := range f.items {
		if !item.IsActive() {
			continue
		}

		item.Update()
	}
}

// Draw implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) Draw(screen S) {
	for _, item := range f.items {
		if !item.IsActive() {
			continue
		}

		item.Draw(screen)
	}
}

// Dispose implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) Dispose() {
	for _, item := range f.items {
		item.Dispose()
	}
}
