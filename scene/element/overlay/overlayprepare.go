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
func (o *overlayImpl[S, T]) Prepare() {
	props := o.Props()

	o.attrs = scene.Attrs{}
	ct := &o.attrs

	minWidth := 0.0
	minHeight := 0.0

	for _, item := range o.items {
		if item.IsOff() {
			item.Exclude()

			continue
		}

		item.Prepare()

		it := item.Attrs()

		minWidth = max(minWidth, basic.FloorPoz(it.MinWidth.Or(0)))
		minHeight = max(minHeight, basic.FloorPoz(it.MinHeight.Or(0)))
	}

	minWidth = max(minWidth, basic.FloorPoz(props.MinWidth.Or(0)))
	minHeight = max(minHeight, basic.FloorPoz(props.MinHeight.Or(0)))

	ct.MinWidth = basic.SetOpt(minWidth)
	ct.MinHeight = basic.SetOpt(minHeight)

	if props.MaxWidth.IsSet() {
		maxWidth := basic.FloorPoz(max(minWidth, basic.FloorPoz(props.MaxWidth.Get())))
		ct.MaxWidth = basic.SetOpt(maxWidth)
	}
	if props.MaxHeight.IsSet() {
		maxHeight := basic.FloorPoz(max(minHeight, basic.FloorPoz(props.MaxHeight.Get())))
		ct.MaxHeight = basic.SetOpt(maxHeight)
	}

	if props.PreWidth.IsSet() {
		maxWidth := ct.MaxWidth.Or(basic.PosInf())
		preWidth := basic.FloorPoz(props.PreWidth.Get())
		ct.PreWidth = basic.SetOpt(basic.Clamp(preWidth, minWidth, maxWidth))
	}
	if props.PreHeight.IsSet() {
		maxHeight := ct.MaxHeight.Or(basic.PosInf())
		preHeight := basic.FloorPoz(props.PreHeight.Get())
		ct.PreHeight = basic.SetOpt(basic.Clamp(preHeight, minHeight, maxHeight))
	}
}
