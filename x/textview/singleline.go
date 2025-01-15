package textview

import (
	"unicode"
	"unicode/utf8"

	"golang.org/x/image/font"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// SingleLine text view.
type SingleLine struct {
	fontFace font.Face
	text     string
	width    Fixed
	height   Fixed
}

// NewSingleLine initializes and returns new [SingleLine] text view with given text, drawn using given font face,
// clipped and extended by ellipsis if needed due to given width hint.
//
// NOTE that non-empty text view will contain at least one rune and ellipsis, even if they overflow width hint.
func NewSingleLine(width basic.Float, fontFace font.Face, text string) SingleLine {
	if fontFace == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	t := SingleLine{
		fontFace: fontFace,
	}

	metrics := t.fontFace.Metrics()
	t.height = metrics.Ascent + metrics.Descent

	if text == "" {
		return t
	}

	spaceAdvance, _ := t.fontFace.GlyphAdvance(' ')

	dotAdvance, _ := t.fontFace.GlyphAdvance('.')
	dotKerning := t.fontFace.Kern('.', '.')
	ellipsisAdvance := dotAdvance + (dotAdvance+dotKerning)*2

	var buf []byte
	var safeLen int
	var safeWidth Fixed
	var space bool
	rPrev := rune(-1)
	for i, size := 0, 0; i < len(text); i += size {
		var r rune
		r, size = utf8.DecodeRuneInString(text[i:])

		if r == '\r' || r == '\n' || unicode.IsSpace(r) {
			space = len(buf) > 0

			continue
		}

		advance, _ := t.fontFace.GlyphAdvance(r)
		if space {
			if rPrev >= 0 {
				advance += t.fontFace.Kern(rPrev, ' ')
			}

			advance += spaceAdvance + t.fontFace.Kern(' ', r)
		} else {
			if rPrev >= 0 {
				advance += t.fontFace.Kern(rPrev, r)
			}
		}

		updatedWidth := t.width + advance

		if FixedToFloat(updatedWidth) > width && len(buf) > 0 {
			buf = buf[:safeLen]
			buf = utf8.AppendRune(buf, '.')
			buf = utf8.AppendRune(buf, '.')
			buf = utf8.AppendRune(buf, '.')
			t.width = safeWidth

			break
		}

		if space {
			buf = utf8.AppendRune(buf, ' ')
		}
		buf = utf8.AppendRune(buf, r)
		t.width = updatedWidth

		updatedSafeWidth := t.width + t.fontFace.Kern(r, '.') + ellipsisAdvance
		if safeLen == 0 || FixedToFloat(updatedSafeWidth) <= width {
			safeLen = len(buf)
			safeWidth = updatedSafeWidth
		}

		space = false
		rPrev = r
	}
	t.text = string(buf)

	return t
}

// IsNil reports whether text view is nil.
func (t SingleLine) IsNil() bool {
	return t.fontFace == nil
}

// Width returns width of text.
func (t SingleLine) Width() basic.Float {
	return FixedToFloat(t.width)
}

// Height returns height of text.
func (t SingleLine) Height() basic.Float {
	return FixedToFloat(t.height)
}

// Draw draws text.
func (t SingleLine) Draw(stringDrawer StringDrawer) {
	if stringDrawer == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	if t.IsNil() {
		return
	}

	stringDrawer.DrawString(t.text, 0, FixedToFloat(t.fontFace.Metrics().Ascent), t.fontFace)
}
