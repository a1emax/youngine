package pageset

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/scene"
)

// Region page can be placed inside.
type Region interface {
	scene.Region

	// Arrange arranges region by setting its bounding rectangle.
	Arrange(rect basic.Rect)
}

// regionImpl is the implementation of the [Region] type.
type regionImpl struct {
	rect basic.Rect
}

// NewRegion initializes and returns new [Region].
func NewRegion() Region {
	return &regionImpl{}
}

// Rect implements the [scene.Region] interface.
func (r *regionImpl) Rect() basic.Rect {
	return r.rect
}

// Arrange implements the [Region] interface.
func (r *regionImpl) Arrange(rect basic.Rect) {
	r.rect = rect
}
