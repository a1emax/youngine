package scene

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// Region element can be placed inside.
type Region interface {

	// Rect returns bounding rectangle of region.
	Rect() basic.Rect
}

// RegionFunc is the functional implementation of the [Region] interface.
type RegionFunc func() basic.Rect

// Rect implements the [Region] interface.
func (f RegionFunc) Rect() basic.Rect {
	if f == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return f()
}
