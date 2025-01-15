package textview

import (
	"unicode"
	"unicode/utf8"

	"golang.org/x/image/font"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// MultiLine text view.
type MultiLine struct {
	fontFace font.Face
	lines    []multiLineLine
	width    Fixed
	height   Fixed
}

// multiLineLine is text line.
type multiLineLine struct {
	text          string
	words         []multiLineWord
	spacedWidth   Fixed
	unspacedWidth Fixed
}

// multiLineWord is word inside text line.
type multiLineWord struct {
	low, high int
	width     Fixed
}

// NewMultiLine initializes and returns new [MultiLine] text view with given text, drawn using given font face,
// broken down into lines by words considering given width hint.
//
// NOTE that each non-empty line will contain at least one rune, even if it overflows width hint.
func NewMultiLine(width basic.Float, fontFace font.Face, text string) MultiLine {
	if fontFace == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	t := MultiLine{
		fontFace: fontFace,
	}

	spaceAdvance, _ := t.fontFace.GlyphAdvance(' ')

	var line multiLineLine
	var lineBuf []byte
	var lineTailKerning Fixed
	newLine := func() {
		line.text = string(lineBuf)
		t.lines = append(t.lines, line)
		t.width = max(t.width, line.spacedWidth)

		line = multiLineLine{}
		lineBuf = lineBuf[:0]

		lineTailKerning = 0
	}

	rPrev := rune(-1)

	var wordBuf []byte
	var wordWidth Fixed
	var wordHeadKerning Fixed
	newWord := func() {
		if len(wordBuf) == 0 {
			return
		}

		if len(lineBuf) > 0 {
			lineBuf = utf8.AppendRune(lineBuf, ' ')

			line.spacedWidth += lineTailKerning
			line.spacedWidth += spaceAdvance
			line.spacedWidth += wordHeadKerning
		}

		i := len(lineBuf)
		lineBuf = append(lineBuf, wordBuf...)
		line.words = append(line.words, multiLineWord{i, i + len(wordBuf), wordWidth})

		line.spacedWidth += wordWidth
		line.unspacedWidth += wordWidth
		lineTailKerning = t.fontFace.Kern(rPrev, ' ')

		wordBuf = wordBuf[:0]
		wordWidth = 0
		wordHeadKerning = 0

		rPrev = rune(-1)
	}

	for i, size := 0, 0; i < len(text); i += size {
		var r rune
		r, size = utf8.DecodeRuneInString(text[i:])

		switch r {
		case '\r':
			if iNext := i + size; iNext < len(text) {
				if rNext, _ := utf8.DecodeRuneInString(text[iNext:]); rNext == '\n' {
					continue
				}
			}
		case '\n':
			newWord()
			newLine()

			continue
		}

		if unicode.IsSpace(r) {
			newWord()

			continue
		}

		wordAdvance, _ := t.fontFace.GlyphAdvance(r)

		var updatedWordWidth Fixed
		var updatedWordHeadKerning Fixed
		if len(wordBuf) > 0 {
			updatedWordWidth = wordWidth + t.fontFace.Kern(rPrev, r) + wordAdvance
			updatedWordHeadKerning = wordHeadKerning
		} else {
			updatedWordWidth = wordAdvance
			if len(lineBuf) > 0 {
				updatedWordHeadKerning = t.fontFace.Kern(' ', r)
			} else {
				updatedWordHeadKerning = 0
			}
		}

		lineAdvance := updatedWordHeadKerning + updatedWordWidth
		if len(lineBuf) > 0 {
			lineAdvance += lineTailKerning + spaceAdvance
		}

		if FixedToFloat(line.spacedWidth+lineAdvance) > width && (len(lineBuf) > 0 || len(wordBuf) > 0) {
			if len(lineBuf) == 0 {
				newWord()
				updatedWordWidth = wordAdvance
			}

			newLine()
			updatedWordHeadKerning = 0
		}

		wordBuf = utf8.AppendRune(wordBuf, r)
		wordWidth = updatedWordWidth
		wordHeadKerning = updatedWordHeadKerning

		rPrev = r
	}
	newWord()
	newLine()

	if n := len(t.lines); n > 0 {
		metrics := t.fontFace.Metrics()
		t.height = metrics.Ascent + metrics.Height*Fixed(n-1) + metrics.Descent
	}

	return t
}

// IsNil reports whether text view is nil.
func (t MultiLine) IsNil() bool {
	return t.fontFace == nil
}

// Width returns width of the widest line.
func (t MultiLine) Width() basic.Float {
	return FixedToFloat(t.width)
}

// Height returns total height of all lines.
func (t MultiLine) Height() basic.Float {
	return FixedToFloat(t.height)
}

// Draw draws text part starting from given offset considering given height hint.
//
// NOTE that line will be drawn if its TOP (can be negative) does not overflow height hint.
func (t MultiLine) Draw(offset, height basic.Float, align Align, stringDrawer StringDrawer) {
	if stringDrawer == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	if t.IsNil() {
		return
	}

	metrics := t.fontFace.Metrics()

	var i int
	var y Fixed

	if offset <= 0 {
		i = 0
		y = -FloatToFixed(offset) + metrics.Ascent
	} else {
		baselineStep := FixedToFloat(metrics.Height)
		ascent := FixedToFloat(metrics.Ascent)
		descent := FixedToFloat(metrics.Descent)

		firstIndex := basic.FloorPoz((offset - ascent) / baselineStep)
		firstBaseline := ascent + firstIndex*baselineStep

		if offset-firstBaseline > descent {
			firstIndex += 1
			firstBaseline += baselineStep
		}

		i = int(firstIndex)
		y = -FloatToFixed(offset - firstBaseline)
	}

	for ; i < len(t.lines); i++ {
		if FixedToFloat(y-metrics.Ascent) >= height {
			break
		}

		line := t.lines[i]

		if line.spacedWidth < t.width && !align.IsDefault() {
			switch align {
			case AlignRight:
				x := t.width - line.spacedWidth
				stringDrawer.DrawString(line.text, FixedToFloat(x), FixedToFloat(y), t.fontFace)
			case AlignCenter:
				x := (t.width - line.spacedWidth) / 2
				stringDrawer.DrawString(line.text, FixedToFloat(x), FixedToFloat(y), t.fontFace)
			case AlignJustify:
				var space Fixed
				if len(line.words) > 1 {
					space = (t.width - line.unspacedWidth) / Fixed(len(line.words)-1)
				}

				var x Fixed
				for _, word := range line.words {
					s := line.text[word.low:word.high]
					stringDrawer.DrawString(s, FixedToFloat(x), FixedToFloat(y), t.fontFace)
					x += word.width + space
				}
			}
		} else {
			stringDrawer.DrawString(line.text, 0, FixedToFloat(y), t.fontFace)
		}

		y += metrics.Height
	}
}
