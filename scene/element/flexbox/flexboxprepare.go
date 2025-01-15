package flexbox

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/scene"
)

// Implementation peculiarities:
//   - Default minimum main size is sum of minimum main sizes of items.
//   - Default minimum cross size is minimum cross size of the biggest item.
//   - Both default maximum sizes are indefinite.
//   - Both default preliminary sizes are indefinite.

// Prepare implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) Prepare() {
	f.outline = scene.Outline{}
	ct := &f.outline

	minWidth := 0.0
	minHeight := 0.0

	for _, item := range f.items {
		if !item.IsActive() {
			item.Exclude()

			continue
		}

		item.Prepare()

		it := item.Outline()

		switch f.state.Direction {
		default: // DirectionRow or invalid:
			minWidth += basic.FloorPoz(it.MinWidth.Or(0))
			minHeight = max(minHeight, basic.FloorPoz(it.MinHeight.Or(0)))
		case DirectionColumn:
			minWidth = max(minWidth, basic.FloorPoz(it.MinWidth.Or(0)))
			minHeight += basic.FloorPoz(it.MinHeight.Or(0))
		}
	}

	minWidth = max(minWidth, basic.FloorPoz(f.state.MinWidth.Or(0)))
	minHeight = max(minHeight, basic.FloorPoz(f.state.MinHeight.Or(0)))

	ct.MinWidth = basic.SetOpt(minWidth)
	ct.MinHeight = basic.SetOpt(minHeight)

	if f.state.MaxWidth.IsSet() {
		maxWidth := max(minWidth, basic.FloorPoz(f.state.MaxWidth.Get()))
		ct.MaxWidth = basic.SetOpt(maxWidth)
	}
	if f.state.MaxHeight.IsSet() {
		maxHeight := max(minHeight, basic.FloorPoz(f.state.MaxHeight.Get()))
		ct.MaxHeight = basic.SetOpt(maxHeight)
	}

	if f.state.PreWidth.IsSet() {
		maxWidth := ct.MaxWidth.Or(basic.PosInf())
		preWidth := basic.FloorPoz(f.state.PreWidth.Get())
		ct.PreWidth = basic.SetOpt(basic.Clamp(preWidth, minWidth, maxWidth))
	}
	if f.state.PreHeight.IsSet() {
		maxHeight := ct.MaxHeight.Or(basic.PosInf())
		preHeight := basic.FloorPoz(f.state.PreHeight.Get())
		ct.PreHeight = basic.SetOpt(basic.Clamp(preHeight, minHeight, maxHeight))
	}
}
