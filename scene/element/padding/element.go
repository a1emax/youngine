package padding

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/scene"
)

// Element places on screen of type S.
type Element[S any] interface {
	scene.Element[S, Trait]
}

// Trait of element.
type Trait struct {

	// Bottom specifies horizontal element offset from container's bottom edge.
	Bottom basic.Opt[basic.Float]

	// Left specifies horizontal element offset from container's left edge.
	Left basic.Opt[basic.Float]

	// Right specifies horizontal element offset from container's right edge.
	Right basic.Opt[basic.Float]

	// Top specifies horizontal element offset from container's top edge.
	Top basic.Opt[basic.Float]
}

// Func returns this trait.
func (t Trait) Func(Trait) Trait {
	return t
}
