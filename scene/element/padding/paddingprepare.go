package padding

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/scene"
)

// Prepare implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) Prepare() {
	p.outline = scene.Outline{}
	ct := &p.outline

	p.element.Prepare()

	state := p.element.Region().State()
	left := basic.FloorPoz(state.Left.Or(0))
	top := basic.FloorPoz(state.Top.Or(0))
	right := basic.FloorPoz(state.Right.Or(0))
	bottom := basic.FloorPoz(state.Bottom.Or(0))

	el := p.element.Outline()

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
