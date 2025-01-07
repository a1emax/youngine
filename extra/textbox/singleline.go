package textbox

import (
	"unicode"
	"unicode/utf8"

	"golang.org/x/image/font"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// SingleLine text box.
type SingleLine struct {
	fontFace font.Face
	text     string
	width    Fixed
	height   Fixed
}

// NewSingleLine initializes and returns new [SingleLine] text box with given text, drawn using given font face,
// clipped and extended by ellipsis if needed due to given width hint.
//
// NOTE that non-empty text box will contain at least one rune and ellipsis, even if they overflow width hint.
func NewSingleLine(width basic.Float, fontFace font.Face, text string) SingleLine {
	if fontFace == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	b := SingleLine{
		fontFace: fontFace,
	}

	metrics := b.fontFace.Metrics()
	b.height = metrics.Ascent + metrics.Descent

	if text == "" {
		return b
	}

	spaceAdvance, _ := b.fontFace.GlyphAdvance(' ')

	dotAdvance, _ := b.fontFace.GlyphAdvance('.')
	dotKerning := b.fontFace.Kern('.', '.')
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

		advance, _ := b.fontFace.GlyphAdvance(r)
		if space {
			if rPrev >= 0 {
				advance += b.fontFace.Kern(rPrev, ' ')
			}

			advance += spaceAdvance + b.fontFace.Kern(' ', r)
		} else {
			if rPrev >= 0 {
				advance += b.fontFace.Kern(rPrev, r)
			}
		}

		updatedWidth := b.width + advance

		if FixedToFloat(updatedWidth) > width && len(buf) > 0 {
			buf = buf[:safeLen]
			buf = utf8.AppendRune(buf, '.')
			buf = utf8.AppendRune(buf, '.')
			buf = utf8.AppendRune(buf, '.')
			b.width = safeWidth

			break
		}

		if space {
			buf = utf8.AppendRune(buf, ' ')
		}
		buf = utf8.AppendRune(buf, r)
		b.width = updatedWidth

		updatedSafeWidth := b.width + b.fontFace.Kern(r, '.') + ellipsisAdvance
		if safeLen == 0 || FixedToFloat(updatedSafeWidth) <= width {
			safeLen = len(buf)
			safeWidth = updatedSafeWidth
		}

		space = false
		rPrev = r
	}
	b.text = string(buf)

	return b
}

// IsNil reports whether text box is nil.
func (b SingleLine) IsNil() bool {
	return b.fontFace == nil
}

// Width returns width of text.
func (b SingleLine) Width() basic.Float {
	return FixedToFloat(b.width)
}

// Height returns height of text.
func (b SingleLine) Height() basic.Float {
	return FixedToFloat(b.height)
}

// Draw draws text.
func (b SingleLine) Draw(stringDrawer StringDrawer) {
	if stringDrawer == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	if b.IsNil() {
		return
	}

	stringDrawer.DrawString(b.text, 0, FixedToFloat(b.fontFace.Metrics().Ascent), b.fontFace)
}
