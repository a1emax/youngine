package padding

import (
	"math"

	"github.com/a1emax/youngine/basic"
)

// containerLayout represents container for internal purposes.
type containerLayout struct {
	xOffset basic.Float
	xSize   basic.Float

	yOffset basic.Float
	ySize   basic.Float
}

// elementLayout represents element for internal purposes.
type elementLayout struct {
	x elementAxis
	y elementAxis
}

// Arrange implements the [scene.Element] interface.
func (p *paddingImpl[S, R]) Arrange() {
	ct := &p.containerLayout

	r := p.region.Rect()
	ct.xOffset = math.Floor(r.Left())
	ct.xSize = basic.FloorPoz(r.Width())
	ct.yOffset = math.Floor(r.Top())
	ct.ySize = basic.FloorPoz(r.Height())

	state := p.element.Region().State()
	outline := p.element.Outline()

	el := &p.elementLayout

	el.x.init(state.Left, state.Right, outline.MinWidth, outline.MaxWidth, ct.xSize)
	el.x.calc()

	el.y.init(state.Top, state.Bottom, outline.MinHeight, outline.MaxHeight, ct.ySize)
	el.y.calc()

	p.element.Region().Arrange(basic.Rect{
		Min: basic.Vec2{
			ct.xOffset + el.x.offset,
			ct.yOffset + el.y.offset,
		},
		Size: basic.Vec2{
			el.x.size.final,
			el.y.size.final,
		},
	})

	p.element.Arrange()
}

// elementAxis contains element properties for one axis (X or Y).
type elementAxis struct {
	before basic.Float
	after  basic.Float

	offset basic.Float
	size   struct {
		min   basic.Float
		max   basic.Float
		outer basic.Float
		base  basic.Float
		final basic.Float
	}
}

// init sets initial properties.
func (a *elementAxis) init(before, after, minSize, maxSize basic.Opt[basic.Float], outerSize basic.Float) {
	a.before = basic.FloorPoz(before.Or(0))
	a.after = basic.FloorPoz(after.Or(0))

	a.size.min = basic.FloorPoz(minSize.Or(0))
	a.size.max = basic.FloorPoz(maxSize.Or(basic.PosInf()))
	a.size.max = max(a.size.min, a.size.max)

	a.size.outer = outerSize
}

// calc sets calculated properties.
func (a *elementAxis) calc() {
	a.offset = a.before
	a.size.base = a.size.outer - a.offset - a.after
	a.size.final = basic.Clamp(a.size.base, a.size.min, a.size.max)
}
