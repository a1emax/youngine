package overlay

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

// itemLayout represents item for internal purposes.
type itemLayout struct {
	index int

	x itemAxis
	y itemAxis
}

// Arrange implements the [scene.Element] interface.
func (o *overlayImpl[S, T]) Arrange(bbox basic.Rect) {
	ct := &o.containerLayout

	ct.xOffset = math.Floor(bbox.Left())
	ct.xSize = basic.FloorPoz(bbox.Width())
	ct.yOffset = math.Floor(bbox.Top())
	ct.ySize = basic.FloorPoz(bbox.Height())

	o.itemLayouts = o.itemLayouts[:0]
	for i, item := range o.items {
		if item.IsOff() {
			continue
		}

		trait := item.Trait()
		attrs := item.Attrs()

		o.itemLayouts = append(o.itemLayouts, itemLayout{})
		it := &o.itemLayouts[len(o.itemLayouts)-1]
		it.index = i

		it.x.init(trait.Left, trait.Right, attrs.MinWidth, attrs.MaxWidth, attrs.PreWidth, ct.xSize)
		it.x.calc()

		it.y.init(trait.Top, trait.Bottom, attrs.MinHeight, attrs.MaxHeight, attrs.PreHeight, ct.ySize)
		it.y.calc()

		item.Arrange(basic.Rect{
			Min: basic.Vec2{
				ct.xOffset + it.x.offset,
				ct.yOffset + it.y.offset,
			},
			Size: basic.Vec2{
				it.x.size.final,
				it.y.size.final,
			},
		})
	}
}

// itemAxis contains item properties for one axis (X or Y).
type itemAxis struct {
	before basic.Float
	after  basic.Float

	offset basic.Float
	size   struct {
		min   basic.Float
		max   basic.Float
		pre   basic.Opt[basic.Float]
		outer basic.Float
		base  basic.Float
		final basic.Float
	}
}

// init sets initial properties.
func (a *itemAxis) init(before, after, minSize, maxSize, preSize basic.Opt[basic.Float], outerSize basic.Float) {
	a.before = basic.Poz(before.Or(1))
	a.after = basic.Poz(after.Or(1))

	a.size.min = basic.FloorPoz(minSize.Or(0))
	a.size.max = basic.FloorPoz(maxSize.Or(basic.PosInf()))
	a.size.max = max(a.size.min, a.size.max)

	a.size.pre = preSize
	a.size.outer = outerSize
}

// calc sets calculated properties.
func (a *itemAxis) calc() {
	if a.size.pre.IsSet() {
		a.size.base = basic.FloorPoz(a.size.pre.Get())
		a.size.base = min(a.size.base, a.size.outer)
	} else {
		a.size.base = a.size.outer
	}

	a.size.final = basic.Clamp(a.size.base, a.size.min, a.size.max)

	remainingFreeSpace := a.size.outer - a.size.final

	ratio := a.before / (a.before + a.after)
	a.offset = math.Floor(remainingFreeSpace * ratio)
}
