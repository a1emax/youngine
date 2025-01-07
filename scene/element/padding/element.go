package padding

import (
	"github.com/a1emax/youngine/scene"
)

// Element places on screen of type S.
type Element[S any] interface {
	scene.Element[S, Region]
}
