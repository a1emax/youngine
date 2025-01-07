package flexbox

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// Region item can be placed inside.
type Region interface {
	scene.Region

	// State returns current state of region.
	State() RegionState

	// Refresh refreshes region by recalculating its state.
	Refresh()

	// Arrange arranges region by setting its bounding rectangle.
	Arrange(rect basic.Rect)
}

// RegionConfig configures [Region].
type RegionConfig struct {

	// StateFunc accepts current state and returns new one.
	StateFunc func(state RegionState) RegionState
}

// RegionState is changeable state of [Region].
type RegionState struct {

	// Basis is initial main size of item.
	Basis basic.Opt[basic.Float]

	// Grow factor of item, which specifies how much of container's remaining space should be assigned
	// to item's main size. If is not set, 1 is used.
	Grow basic.Opt[basic.Float]

	// Shrink factor of item. If size of all items is larger than container, items shrink to fit according
	// to this factor. If is not set, 1 is used.
	Shrink basic.Opt[basic.Float]

	// AlignSelf specifies alignment of item on container's cross axis.
	AlignSelf basic.Opt[Align]
}

// regionImpl is the implementation of the [Region] type.
type regionImpl struct {
	RegionConfig

	rect  basic.Rect
	state RegionState
}

// NewRegion initializes and returns new [Region].
func NewRegion(config RegionConfig) Region {
	if config.StateFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &regionImpl{
		RegionConfig: config,
	}
}

// Rect implements the [scene.Region] interface.
func (r *regionImpl) Rect() basic.Rect {
	return r.rect
}

// State implements the [Region] interface.
func (r *regionImpl) State() RegionState {
	return r.state
}

// Refresh implements the [Region] interface.
func (r *regionImpl) Refresh() {
	r.state = r.StateFunc(r.state)
}

// Arrange implements the [Region] interface.
func (r *regionImpl) Arrange(rect basic.Rect) {
	r.rect = rect
}
