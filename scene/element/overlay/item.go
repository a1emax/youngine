package overlay

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/scene"
)

// Item displayed on screen of type S.
type Item[S any] interface {
	scene.Element[S, Trait]
}

// Trait of item.
type Trait struct {

	// Bottom space factor, which specifies how much of container's remaining horizontal space should be
	// to bottom of item. If is not set, 1 is used.
	Bottom basic.Opt[basic.Float]

	// Left space factor, which specifies how much of container's remaining horizontal space should be
	// to left of item. If is not set, 1 is used.
	Left basic.Opt[basic.Float]

	// Right space factor, which specifies how much of container's remaining horizontal space should be
	// to right of item. If is not set, 1 is used.
	Right basic.Opt[basic.Float]

	// Top space factor, which specifies how much of container's remaining horizontal space should be
	// to top of item. If is not set, 1 is used.
	Top basic.Opt[basic.Float]
}

// Func returns new [scene.TraitFunc] that returns this trait.
func (t Trait) Func(Trait) Trait {
	return t
}
