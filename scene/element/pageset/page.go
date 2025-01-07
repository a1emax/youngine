package pageset

import (
	"github.com/a1emax/youngine/scene"
)

// Page placed on screen of type S.
type Page[S any] interface {
	scene.Element[S, Region]
}
