package overlay

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// Overlay placed on screen of type S inside region of type R.
type Overlay[S any, R scene.Region] interface {
	scene.Element[S, R]
}

// Config configures [Overlay].
type Config struct {

	// StateFunc accepts current state and returns new one.
	StateFunc func(state State) State
}

// State is changeable state of [Overlay].
type State struct {

	// IsInactive specifies whether element is inactive.
	IsInactive bool

	// Outline specifies base outline of element.
	scene.Outline
}

// overlayImpl is the implementation of the [Overlay] interface.
type overlayImpl[S any, R scene.Region] struct {
	scene.BaseElement[S, R]
	Config

	region R
	state  State
	items  []Item[S]

	outline         scene.Outline
	containerLayout containerLayout
	itemLayouts     []itemLayout
}

// New initializes and returns new [Overlay].
func New[S any, R scene.Region](region R, config Config, items ...Item[S]) Overlay[S, R] {
	if config.StateFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	o := &overlayImpl[S, R]{
		Config: config,

		region: region,
	}

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

// Region implements the [scene.Element] interface.
func (o *overlayImpl[S, R]) Region() R {
	return o.region
}

// IsActive implements the [scene.Element] interface.
func (o *overlayImpl[S, R]) IsActive() bool {
	return !o.state.IsInactive
}

// Outline implements the [scene.Element] interface.
func (o *overlayImpl[S, R]) Outline() scene.Outline {
	return o.outline
}

// Refresh implements the [scene.Element] interface.
func (o *overlayImpl[S, R]) Refresh() {
	o.state = o.StateFunc(o.state)
	o.outline = scene.Outline{}

	for _, item := range o.items {
		item.Region().Refresh()
		item.Refresh()
	}
}

// Exclude implements the [scene.Element] interface.
func (o *overlayImpl[S, R]) Exclude() {
	for _, item := range o.items {
		item.Exclude()
	}
}

// Actuate implements the [scene.Element] interface.
func (o *overlayImpl[S, R]) Actuate() {
	for i := len(o.items) - 1; i >= 0; i-- {
		item := o.items[i]

		if item.IsActive() {
			item.Actuate()
		} else {
			item.Inhibit()
		}
	}
}

// Inhibit implements the [scene.Element] interface.
func (o *overlayImpl[S, R]) Inhibit() {
	for i := len(o.items) - 1; i >= 0; i-- {
		item := o.items[i]

		item.Inhibit()
	}
}

// Update implements the [scene.Element] interface.
func (o *overlayImpl[S, R]) Update() {
	for _, item := range o.items {
		if !item.IsActive() {
			continue
		}

		item.Update()
	}
}

// Draw implements the [scene.Element] interface.
func (o *overlayImpl[S, R]) Draw(screen S) {
	for _, item := range o.items {
		if !item.IsActive() {
			continue
		}

		item.Draw(screen)
	}
}

// Dispose implements the [scene.Element] interface.
func (o *overlayImpl[S, R]) Dispose() {
	for _, item := range o.items {
		item.Dispose()
	}
}
