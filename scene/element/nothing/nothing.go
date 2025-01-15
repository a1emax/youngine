package nothing

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// Nothing placed on screen of type S inside region of type R.
type Nothing[S any, R scene.Region] interface {
	scene.Element[S, R]
}

// Config configures [Nothing].
type Config struct {

	// StateFunc accepts current state and returns new one.
	StateFunc func(state State) State
}

// State is changeable state of [Nothing].
type State struct {

	// IsInactive specifies whether element is inactive.
	IsInactive bool

	// Outline of element.
	scene.Outline
}

// nothingImpl is the implementation of the [Nothing] interface.
type nothingImpl[S any, R scene.Region] struct {
	scene.BaseElement[S, R]
	Config

	region R
	state  State
}

// New initializes and returns new [Nothing].
func New[S any, R scene.Region](region R, config Config) Nothing[S, R] {
	if config.StateFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &nothingImpl[S, R]{
		Config: config,

		region: region,
	}
}

// Region implements the [scene.Element] interface.
func (n *nothingImpl[S, R]) Region() R {
	return n.region
}

// IsActive implements the [scene.Element] interface.
func (n *nothingImpl[S, R]) IsActive() bool {
	return !n.state.IsInactive
}

// Outline implements the [scene.Element] interface.
func (n *nothingImpl[S, R]) Outline() scene.Outline {
	return n.state.Outline
}

// Refresh implements the [scene.Element] interface.
func (n *nothingImpl[S, R]) Refresh() {
	n.state = n.StateFunc(n.state)
}
