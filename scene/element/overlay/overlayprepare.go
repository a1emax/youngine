package overlay

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/scene"
)

// Implementation peculiarities:
//   - Default minimum width is minimum width of the widest item.
//   - Default minimum height is minimum height of the highest item.
//   - Both default maximum sizes are indefinite.
//   - Both default preliminary sizes are indefinite.

// Prepare implements the [scene.Element] interface.
func (o *overlayImpl[S, R]) Prepare() {
	o.outline = scene.Outline{}
	ct := &o.outline

	minWidth := 0.0
	minHeight := 0.0

	for _, item := range o.items {
		if !item.IsActive() {
			item.Exclude()

			continue
		}

		item.Prepare()

		it := item.Outline()

		minWidth = max(minWidth, basic.FloorPoz(it.MinWidth.Or(0)))
		minHeight = max(minHeight, basic.FloorPoz(it.MinHeight.Or(0)))
	}

	minWidth = max(minWidth, basic.FloorPoz(o.state.MinWidth.Or(0)))
	minHeight = max(minHeight, basic.FloorPoz(o.state.MinHeight.Or(0)))

	ct.MinWidth = basic.SetOpt(minWidth)
	ct.MinHeight = basic.SetOpt(minHeight)

	if o.state.MaxWidth.IsSet() {
		maxWidth := basic.FloorPoz(max(minWidth, basic.FloorPoz(o.state.MaxWidth.Get())))
		ct.MaxWidth = basic.SetOpt(maxWidth)
	}
	if o.state.MaxHeight.IsSet() {
		maxHeight := basic.FloorPoz(max(minHeight, basic.FloorPoz(o.state.MaxHeight.Get())))
		ct.MaxHeight = basic.SetOpt(maxHeight)
	}

	if o.state.PreWidth.IsSet() {
		maxWidth := ct.MaxWidth.Or(basic.PosInf())
		preWidth := basic.FloorPoz(o.state.PreWidth.Get())
		ct.PreWidth = basic.SetOpt(basic.Clamp(preWidth, minWidth, maxWidth))
	}
	if o.state.PreHeight.IsSet() {
		maxHeight := ct.MaxHeight.Or(basic.PosInf())
		preHeight := basic.FloorPoz(o.state.PreHeight.Get())
		ct.PreHeight = basic.SetOpt(basic.Clamp(preHeight, minHeight, maxHeight))
	}
}
