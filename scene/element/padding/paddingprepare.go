package padding

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/scene"
)

// Prepare implements the [scene.Element] interface.
func (p *paddingImpl[S, T]) Prepare() {
	p.attrs = scene.Attrs{}
	ct := &p.attrs

	p.element.Prepare()

	trait := p.element.Trait()
	left := basic.FloorPoz(trait.Left.Or(0))
	top := basic.FloorPoz(trait.Top.Or(0))
	right := basic.FloorPoz(trait.Right.Or(0))
	bottom := basic.FloorPoz(trait.Bottom.Or(0))

	el := p.element.Attrs()

	minWidth := left + basic.FloorPoz(el.MinWidth.Or(0)) + right
	minHeight := top + basic.FloorPoz(el.MinHeight.Or(0)) + bottom

	ct.MinWidth = basic.SetOpt(minWidth)
	ct.MinHeight = basic.SetOpt(minHeight)

	if el.MaxWidth.IsSet() {
		maxWidth := max(minWidth, left+basic.FloorPoz(el.MaxWidth.Or(0))+right)
		ct.MaxWidth = basic.SetOpt(maxWidth)
	}
	if el.MaxHeight.IsSet() {
		maxHeight := max(minHeight, top+basic.FloorPoz(el.MaxHeight.Or(0))+bottom)
		ct.MaxHeight = basic.SetOpt(maxHeight)
	}

	if el.PreWidth.IsSet() {
		maxWidth := ct.MaxWidth.Or(basic.PosInf())
		preWidth := left + el.PreWidth.Or(0) + right
		ct.PreWidth = basic.SetOpt(basic.Clamp(preWidth, minWidth, maxWidth))
	}
	if el.PreHeight.IsSet() {
		maxHeight := ct.MaxHeight.Or(basic.PosInf())
		preHeight := top + el.PreHeight.Or(0) + bottom
		ct.PreHeight = basic.SetOpt(basic.Clamp(preHeight, minHeight, maxHeight))
	}
}
