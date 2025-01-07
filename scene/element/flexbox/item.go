package flexbox

import (
	"github.com/a1emax/youngine/scene"
)

// Item placed on screen of type S.
type Item[S any] interface {
	scene.Element[S, Region]
}
