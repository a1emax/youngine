package turnoff

import (
	"github.com/a1emax/youngine/scene"
)

// Element displayed on screen of type S.
type Element[S any] interface {
	scene.Element[S, Trait]
}

// Trait of element.
type Trait struct {
}

// Func returns this trait.
func (t Trait) Func(Trait) Trait {
	return t
}
