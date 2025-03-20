package pageset

import (
	"github.com/a1emax/youngine/scene"
)

// Page displayed on screen of type S.
type Page[S any] interface {
	scene.Element[S, Trait]
}

// Trait of page.
type Trait struct {
}

// Func returns this trait.
func (t Trait) Func(Trait) Trait {
	return t
}
