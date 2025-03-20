package textview

import (
	"math"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"golang.org/x/image/font"
)

// DrawFunc draws s using given font face starting from ORIGIN at (x, y).
type DrawFunc func(s string, fontFace font.Face, x, y basic.Float)

// Draw draws view considering given outer size.
func (v TextView) Draw(outerSize basic.Vec2, justifyLines, justifyWords Justify, f DrawFunc) {
	if f == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	if v.IsNil() {
		return
	}

	metrics := v.fontFace.Metrics()

	if !justifyLines.IsValid() {
		justifyLines = JustifyStart
	}
	if FixedToFloat(v.height) > outerSize.Y() {
		switch justifyLines {
		case JustifySpaceBetween:
			justifyLines = JustifyStart
		case JustifySpaceAround, JustifySpaceEvenly:
			justifyLines = JustifyCenter
		}
	}
	if justifyLines <= JustifyEnd {
		var y Fixed
		switch justifyLines {
		case JustifyStart:
			y = 0
		case JustifyCenter:
			y = FloatToFixed(math.Floor(outerSize.Y()-FixedToFloat(v.height)) / 2)
		case JustifyEnd:
			y = FloatToFixed(math.Floor(outerSize.Y() - FixedToFloat(v.height)))
		}

		y += metrics.Ascent
		for i := range v.lines {
			v.drawLine(i, outerSize.X(), FixedToFloat(y), justifyWords, f)

			y += metrics.Height
		}
	} else {
		lineHeight := math.Floor(FixedToFloat(metrics.Ascent + metrics.Descent))

		n := len(v.lines)
		remainingFreeSpace := outerSize.Y() - lineHeight*basic.Float(n)

		var spacing, offset basic.Float
		switch justifyLines {
		case JustifySpaceBetween:
			if n > 1 {
				spacing = math.Floor(remainingFreeSpace / basic.Float(n-1))
			} else {
				spacing = 0
			}
			offset = 0
		case JustifySpaceAround:
			if n > 0 {
				spacing = math.Floor(remainingFreeSpace / basic.Float(n))
			} else {
				spacing = 0
			}
			offset = math.Floor(spacing / 2)
		case JustifySpaceEvenly:
			spacing = math.Floor(remainingFreeSpace / basic.Float(n+1))
			offset = spacing
		}

		for i := range v.lines {
			y := FloatToFixed(offset) + metrics.Ascent
			offset += spacing + lineHeight

			v.drawLine(i, outerSize.X(), FixedToFloat(y), justifyWords, f)
		}
	}
}

// DrawScrollable draws part of view considering given outer size, starting from given vertical scroll offset.
func (v TextView) DrawScrollable(outerSize basic.Vec2, scrollOffset basic.Float, justifyWords Justify, f DrawFunc) {
	if f == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	if v.IsNil() {
		return
	}

	metrics := v.fontFace.Metrics()

	var i int
	var y Fixed

	if scrollOffset <= 0 {
		i = 0
		y = -FloatToFixed(scrollOffset) + metrics.Ascent
	} else {
		baselineStep := FixedToFloat(metrics.Height)
		ascent := FixedToFloat(metrics.Ascent)
		descent := FixedToFloat(metrics.Descent)

		firstIndex := basic.FloorPoz((scrollOffset - ascent) / baselineStep)
		firstBaseline := ascent + firstIndex*baselineStep

		if scrollOffset-firstBaseline > descent {
			firstIndex += 1
			firstBaseline += baselineStep
		}

		i = int(firstIndex)
		y = -FloatToFixed(scrollOffset - firstBaseline)
	}

	for ; i < len(v.lines); i++ {
		if FixedToFloat(y-metrics.Ascent) >= outerSize.Y() {
			break
		}

		v.drawLine(i, outerSize.X(), FixedToFloat(y), justifyWords, f)

		y += metrics.Height
	}
}

// drawLine draws line with given index considering given outer width and baseline position.
func (v TextView) drawLine(i int, outerWidth, y basic.Float, justifyWords Justify, f DrawFunc) {
	line := v.lines[i]

	if !justifyWords.IsValid() {
		justifyWords = JustifyStart
	}
	if FixedToFloat(line.width) > outerWidth {
		switch justifyWords {
		case JustifySpaceBetween:
			justifyWords = JustifyStart
		case JustifySpaceAround, JustifySpaceEvenly:
			justifyWords = JustifyCenter
		}
	}
	if justifyWords <= JustifyEnd {
		var x basic.Float
		switch justifyWords {
		case JustifyStart:
			x = 0
		case JustifyCenter:
			x = math.Floor((outerWidth - FixedToFloat(line.width)) / 2)
		case JustifyEnd:
			x = math.Floor(outerWidth - FixedToFloat(line.width))
		}

		f(line.str, v.fontFace, x, y)
	} else {
		n := len(line.words)
		remainingFreeSpace := outerWidth - FixedToFloat(line.unspacedWidth)

		var spacing, offset basic.Float
		switch justifyWords {
		case JustifySpaceBetween:
			if n > 1 {
				spacing = math.Floor(remainingFreeSpace / basic.Float(n-1))
			} else {
				spacing = 0
			}
			offset = 0
		case JustifySpaceAround:
			if n > 0 {
				spacing = math.Floor(remainingFreeSpace / basic.Float(n))
			} else {
				spacing = 0
			}
			offset = math.Floor(spacing / 2)
		case JustifySpaceEvenly:
			spacing = math.Floor(remainingFreeSpace / basic.Float(n+1))
			offset = spacing
		}

		for _, word := range line.words {
			x := offset
			offset += spacing + math.Floor(FixedToFloat(word.width))

			f(line.str[word.low:word.high], v.fontFace, x, y)
		}
	}
}
