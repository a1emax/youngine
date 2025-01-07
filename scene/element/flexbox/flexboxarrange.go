package flexbox

import (
	"math"

	"github.com/a1emax/youngine/basic"
)

// Flex Layout Algorithm
// https://www.w3.org/TR/css-flexbox-1/#layout-algorithm

// Implementation peculiarities:
//   - Margins, borders and paddings are not supported.
//   - Multi-line containers are not supported.
//   - Intrinsic sizes are not supported.
//   - Fragmentation is not supported.
//   - Container has definite sizes (obtained using region).
//   - Base main size of item that has neither basis nor preliminary main size is considered zero.
//   - An item that has no preliminary cross size is considered stretching.

// containerLayout represents container for internal purposes.
type containerLayout struct {
	mainOffset basic.Float
	mainSize   basic.Float

	crossOffset basic.Float
	crossSize   basic.Float
}

// itemLayout represents item for internal purposes.
type itemLayout struct {
	index int

	basis     basic.Opt[basic.Float]
	grow      basic.Float
	shrink    basic.Float
	alignSelf basic.Opt[Align]

	mainOffset basic.Float
	mainSize   struct {
		min       basic.Float
		max       basic.Float
		pre       basic.Opt[basic.Float]
		base      basic.Float
		hypot     basic.Float
		factor    basic.Float
		violation basic.Float
		isFrozen  bool
		final     basic.Float
	}

	crossOffset basic.Float
	crossSize   struct {
		min          basic.Float
		max          basic.Float
		pre          basic.Opt[basic.Float]
		hypot        basic.Float
		forceStretch bool
		final        basic.Float
	}
}

// Arrange implements the [scene.Element] interface.
func (f *flexboxImpl[S, R]) Arrange() {
	ct := &f.containerLayout

	switch r := f.region.Rect(); f.state.Direction {
	default: // DirectionRow or invalid:
		ct.mainOffset = math.Floor(r.Left())
		ct.mainSize = basic.FloorPoz(r.Width())
		ct.crossOffset = math.Floor(r.Top())
		ct.crossSize = basic.FloorPoz(r.Height())
	case DirectionColumn:
		ct.mainOffset = math.Floor(r.Top())
		ct.mainSize = basic.FloorPoz(r.Height())
		ct.crossOffset = math.Floor(r.Left())
		ct.crossSize = basic.FloorPoz(r.Width())
	}

	var mainHypotSum basic.Float

	f.itemLayouts = f.itemLayouts[:0]
	for i, item := range f.items {
		if !item.IsActive() {
			continue
		}

		state := item.Region().State()
		outline := item.Outline()

		f.itemLayouts = append(f.itemLayouts, itemLayout{})
		it := &f.itemLayouts[len(f.itemLayouts)-1]
		it.index = i

		it.basis = state.Basis
		it.grow = basic.Poz(state.Grow.Or(1))
		it.shrink = basic.Poz(state.Shrink.Or(1))
		it.alignSelf = state.AlignSelf

		switch f.state.Direction {
		default: // DirectionRow or invalid:
			it.mainSize.min = basic.FloorPoz(outline.MinWidth.Or(0))
			it.mainSize.max = basic.FloorPoz(outline.MaxWidth.Or(basic.PosInf()))
			it.mainSize.pre = outline.PreWidth
			it.crossSize.min = basic.FloorPoz(outline.MinHeight.Or(0))
			it.crossSize.max = basic.FloorPoz(outline.MaxHeight.Or(basic.PosInf()))
			it.crossSize.pre = outline.PreHeight
		case DirectionColumn:
			it.mainSize.min = basic.FloorPoz(outline.MinHeight.Or(0))
			it.mainSize.max = basic.FloorPoz(outline.MaxHeight.Or(basic.PosInf()))
			it.mainSize.pre = outline.PreHeight
			it.crossSize.min = basic.FloorPoz(outline.MinWidth.Or(0))
			it.crossSize.max = basic.FloorPoz(outline.MaxWidth.Or(basic.PosInf()))
			it.crossSize.pre = outline.PreWidth
		}

		it.mainSize.max = max(it.mainSize.min, it.mainSize.max)
		it.crossSize.max = max(it.crossSize.min, it.crossSize.max)

		if it.basis.IsSet() {
			it.mainSize.base = basic.FloorPoz(it.basis.Get())
			it.mainSize.base = min(it.mainSize.base, ct.mainSize)
		} else if it.mainSize.pre.IsSet() {
			it.mainSize.base = basic.FloorPoz(it.mainSize.pre.Get())
			it.mainSize.base = min(it.mainSize.base, ct.mainSize)
		} else {
			it.mainSize.base = 0
		}

		it.mainSize.hypot = basic.Clamp(it.mainSize.base, it.mainSize.min, it.mainSize.max)

		mainHypotSum += it.mainSize.hypot

		if it.crossSize.pre.IsSet() {
			it.crossSize.hypot = basic.FloorPoz(it.crossSize.pre.Get())
			it.crossSize.hypot = min(it.crossSize.hypot, ct.crossSize)
		} else {
			it.crossSize.hypot = 0
			it.crossSize.forceStretch = true
		}
	}

	grow := mainHypotSum < ct.mainSize
	shrink := mainHypotSum > ct.mainSize
	// Otherwise all items will be frozen on initial stage.

	initialFreeMainSpace := ct.mainSize

	for i := range f.itemLayouts {
		it := &f.itemLayouts[i]

		switch {
		case grow:
			it.mainSize.factor = it.grow
		case shrink:
			// Scaled flex shrink factor.
			it.mainSize.factor = it.shrink * it.mainSize.base
		}

		switch {
		case it.mainSize.factor == 0:
			it.mainSize.isFrozen = true
		case grow:
			it.mainSize.isFrozen = it.mainSize.base > it.mainSize.hypot
		case shrink:
			it.mainSize.isFrozen = it.mainSize.base < it.mainSize.hypot
		}

		if it.mainSize.isFrozen {
			it.mainSize.final = it.mainSize.hypot
			initialFreeMainSpace -= it.mainSize.final
		} else {
			initialFreeMainSpace -= it.mainSize.base
		}
	}

	var remainingFreeMainSpace basic.Float

	for {
		var frozenCount int
		var factorSum basic.Float

		remainingFreeMainSpace = ct.mainSize

		for i := range f.itemLayouts {
			it := &f.itemLayouts[i]

			if it.mainSize.isFrozen {
				frozenCount++
				remainingFreeMainSpace -= it.mainSize.final
			} else {
				factorSum += it.mainSize.factor
				remainingFreeMainSpace -= it.mainSize.base
			}
		}

		if frozenCount == len(f.itemLayouts) {
			break
		}

		if factorSum < 1 {
			if v := math.Floor(initialFreeMainSpace * factorSum); math.Abs(v) < math.Abs(remainingFreeMainSpace) {
				remainingFreeMainSpace = v
			}
		}

		var totalViolation basic.Float

		for i := range f.itemLayouts {
			it := &f.itemLayouts[i]

			if it.mainSize.isFrozen {
				continue
			}

			if remainingFreeMainSpace > 0 {
				ratio := it.mainSize.factor / factorSum
				switch {
				case grow:
					it.mainSize.final = it.mainSize.base + remainingFreeMainSpace*ratio
				case shrink:
					// Note this may result in negative main size; it will be corrected in next step.
					it.mainSize.final = it.mainSize.base - math.Abs(remainingFreeMainSpace)*ratio
				}
			}

			clampedMainSize := basic.Clamp(it.mainSize.final, it.mainSize.min, it.mainSize.max)
			it.mainSize.violation = clampedMainSize - it.mainSize.final
			it.mainSize.final = math.Floor(clampedMainSize)

			totalViolation += it.mainSize.violation
		}

		for i := range f.itemLayouts {
			it := &f.itemLayouts[i]

			if it.mainSize.isFrozen {
				continue
			}

			switch {
			case totalViolation > 0:
				it.mainSize.isFrozen = it.mainSize.violation > 0
			case totalViolation < 0:
				it.mainSize.isFrozen = it.mainSize.violation < 0
			default:
				it.mainSize.isFrozen = true
			}
		}
	}

	// Remaining free space is non-negative here.
	var mainSpacing basic.Float
	var mainOffset basic.Float
	if n := len(f.itemLayouts); n > 0 {
		switch f.state.JustifyContent {
		default: // JustifyStart or invalid:
			mainSpacing = 0
			mainOffset = 0
		case JustifyCenter:
			mainSpacing = 0
			mainOffset = math.Floor(remainingFreeMainSpace / 2)
		case JustifyEnd:
			mainSpacing = 0
			mainOffset = remainingFreeMainSpace
		case JustifySpaceBetween:
			mainSpacing = math.Floor(remainingFreeMainSpace / basic.Float(n-1))
			mainOffset = 0
		case JustifySpaceAround:
			mainSpacing = math.Floor(remainingFreeMainSpace / basic.Float(n))
			mainOffset = math.Floor(mainSpacing / 2)
		case JustifySpaceEvenly:
			mainSpacing = math.Floor(remainingFreeMainSpace / basic.Float(n+1))
			mainOffset = mainSpacing
		}
	}

	for i := range f.itemLayouts {
		it := &f.itemLayouts[i]

		it.mainOffset = mainOffset
		mainOffset += mainSpacing + it.mainSize.final

		var align Align
		if it.crossSize.forceStretch {
			align = AlignStretch
		} else {
			align = it.alignSelf.Or(f.state.AlignItems)
		}

		var crossSize basic.Float
		if align == AlignStretch {
			crossSize = ct.crossSize
		} else {
			crossSize = it.crossSize.hypot
		}

		it.crossSize.final = basic.Clamp(crossSize, it.crossSize.min, it.crossSize.max)

		switch align {
		default: // AlignStretch, AlignCenter or invalid:
			it.crossOffset = math.Floor((ct.crossSize - it.crossSize.final) / 2)
		case AlignStart:
			it.crossOffset = 0
		case AlignEnd:
			it.crossOffset = ct.crossSize - it.crossSize.final
		}

		item := f.items[it.index]

		switch f.state.Direction {
		default: // DirectionRow or invalid:
			item.Region().Arrange(basic.Rect{
				Min: basic.Vec2{
					ct.mainOffset + it.mainOffset,
					ct.crossOffset + it.crossOffset,
				},
				Size: basic.Vec2{
					it.mainSize.final,
					it.crossSize.final,
				},
			})
		case DirectionColumn:
			item.Region().Arrange(basic.Rect{
				Min: basic.Vec2{
					ct.crossOffset + it.crossOffset,
					ct.mainOffset + it.mainOffset,
				},
				Size: basic.Vec2{
					it.crossSize.final,
					it.mainSize.final,
				},
			})
		}

		item.Arrange()
	}
}
