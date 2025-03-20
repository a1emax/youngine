package flexbox

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

	// AlignSelf specifies alignment of item on container's cross axis.
	AlignSelf basic.Opt[Align]

	// Basis is initial main size of item.
	Basis basic.Opt[basic.Float]

	// Grow factor of item, which specifies how much of container's remaining space should be assigned
	// to item's main size. If is not set, 1 is used.
	Grow basic.Opt[basic.Float]

	// Shrink factor of item. If size of all items is larger than container, items shrink to fit according
	// to this factor. If is not set, 1 is used.
	Shrink basic.Opt[basic.Float]
}

// Func returns this trait.
func (t Trait) Func(Trait) Trait {
	return t
}
