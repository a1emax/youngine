package textview

import (
	"unicode"
	"unicode/utf8"

	"github.com/a1emax/youngine/basic"
	"golang.org/x/image/font"
)

// builder of view.
type builder struct {
	view            TextView
	text            string
	sizeLimit       basic.Vec2
	metrics         font.Metrics
	spaceAdvance    Fixed
	ellipsisAdvance Fixed
	height          Fixed
	line            builderLine
	word            builderWord
	safe            builderSafe
}

// builderLine represents incomplete line.
type builderLine struct {
	buf           []byte
	words         []textViewWord
	width         Fixed
	unspacedWidth Fixed
	tailKerning   Fixed
}

// builderWord represents incomplete word.
type builderWord struct {
	buf         []byte
	lastRune    rune
	width       Fixed
	headKerning Fixed
}

// builderSafe represents safe state of line and word that still allows to extend line with ellipsis.
type builderSafe struct {
	init bool
	line struct {
		bufLen        int
		width         Fixed
		unspacedWidth Fixed
		tailKerning   Fixed
	}
	word struct {
		index         int
		bufLen        int
		advancedWidth Fixed
		headKerning   Fixed
	}
}

// build builds view.
func (b *builder) build() {
	b.metrics = b.view.fontFace.Metrics()

	b.spaceAdvance, _ = b.view.fontFace.GlyphAdvance(' ')

	dotAdvance, _ := b.view.fontFace.GlyphAdvance('.')
	dotKerning := b.view.fontFace.Kern('.', '.')
	b.ellipsisAdvance = dotAdvance + (dotAdvance+dotKerning)*2

	b.height = b.metrics.Ascent

	for i, size := 0, 0; i < len(b.text); i += size {
		var r rune
		r, size = utf8.DecodeRuneInString(b.text[i:])

		if unicode.IsSpace(r) {
			b.flushWord()
			if r == '\n' {
				if b.checkSafe() {
					break
				}

				b.flushLine()
			}

			continue
		}

		lineAdvance, ok := b.appendRune(r)
		if !ok {
			if len(b.line.buf) == 0 {
				b.flushWord()
			}

			if b.checkSafe() {
				break
			}

			b.flushLine()

			size = 0

			continue
		}

		b.updateSafe(lineAdvance)
	}
	b.flushWord()
	b.flushLine()

	b.view.height += b.metrics.Descent
}

// flushLine appends current line to view and resets line state.
func (b *builder) flushLine() {
	line := textViewLine{string(b.line.buf), b.line.words, b.line.width, b.line.unspacedWidth}
	b.line.buf = b.line.buf[:0]
	b.line.words = nil
	b.line.width = 0
	b.line.unspacedWidth = 0
	b.line.tailKerning = 0

	b.view.lines = append(b.view.lines, line)
	b.view.width = max(b.view.width, line.width)
	b.view.height = b.height
	b.height += b.metrics.Height

	b.safe = builderSafe{}
}

// flushWord appends non-empty current word to current line and resets word state.
func (b *builder) flushWord() {
	if len(b.word.buf) == 0 {
		return
	}

	if len(b.line.buf) > 0 {
		b.line.buf = utf8.AppendRune(b.line.buf, ' ')
		b.line.width += b.line.tailKerning + b.spaceAdvance + b.word.headKerning
	}

	i := len(b.line.buf)
	b.line.buf = append(b.line.buf, b.word.buf...)
	b.line.words = append(b.line.words, textViewWord{i, i + len(b.word.buf), b.word.width})

	b.line.width += b.word.width
	b.line.unspacedWidth += b.word.width
	b.line.tailKerning = b.view.fontFace.Kern(b.word.lastRune, ' ')

	b.word.buf = b.word.buf[:0]
	b.word.lastRune = rune(-1)
	b.word.width = 0
	b.word.headKerning = 0
}

// appendRune appends given rune to current word if after that current line will contain one rune
// or will not overflow width limit.
func (b *builder) appendRune(r rune) (lineAdvance Fixed, ok bool) {
	wordAdvance, _ := b.view.fontFace.GlyphAdvance(r)

	var newWordWidth Fixed
	var newWordHeadKerning Fixed
	if len(b.word.buf) > 0 {
		newWordWidth = b.word.width + b.view.fontFace.Kern(b.word.lastRune, r) + wordAdvance
		newWordHeadKerning = b.word.headKerning
	} else {
		newWordWidth = wordAdvance
		if len(b.line.buf) > 0 {
			newWordHeadKerning = b.view.fontFace.Kern(' ', r)
		} else {
			newWordHeadKerning = 0
		}
	}

	lineAdvance = newWordHeadKerning + newWordWidth
	if len(b.line.buf) > 0 {
		lineAdvance += b.line.tailKerning + b.spaceAdvance
	}

	newLineWidth := b.line.width + lineAdvance
	if FixedToFloat(newLineWidth) > b.sizeLimit.X() && (len(b.line.buf) > 0 || len(b.word.buf) > 0) {
		return 0, false
	}

	b.word.buf = utf8.AppendRune(b.word.buf, r)
	b.word.lastRune = r
	b.word.width = newWordWidth
	b.word.headKerning = newWordHeadKerning

	return lineAdvance, true
}

// updateSafe updates safe state.
func (b *builder) updateSafe(lineAdvance Fixed) {
	safeWordAdvance := b.view.fontFace.Kern(b.word.lastRune, '.') + b.ellipsisAdvance
	safeLineWidth := b.line.width + lineAdvance + safeWordAdvance
	if !b.safe.init || FixedToFloat(safeLineWidth) <= b.sizeLimit.X() {
		b.safe.init = true
		b.safe.line.bufLen = len(b.line.buf)
		b.safe.line.width = b.line.width
		b.safe.line.unspacedWidth = b.line.unspacedWidth
		b.safe.line.tailKerning = b.line.tailKerning
		b.safe.word.index = len(b.line.words)
		b.safe.word.bufLen = len(b.word.buf)
		b.safe.word.advancedWidth = b.word.width + safeWordAdvance
		b.safe.word.headKerning = b.word.headKerning
	}
}

// checkSafe, if next line will overflow height limit, restores safe state and then flushes word.
func (b *builder) checkSafe() bool {
	newHeight := b.height + b.metrics.Height + b.metrics.Descent
	if FixedToFloat(newHeight) <= b.sizeLimit.Y() {
		return false
	}

	if b.safe.word.index < len(b.line.words) {
		word := b.line.words[b.safe.word.index]
		b.word.buf = b.line.buf[word.low:word.high]

		b.line.buf = b.line.buf[:b.safe.line.bufLen]
		b.line.words = b.line.words[:b.safe.word.index]
		b.line.width = b.safe.line.width
		b.line.unspacedWidth = b.safe.line.unspacedWidth
		b.line.tailKerning = b.safe.line.tailKerning
	}

	b.word.buf = b.word.buf[:b.safe.word.bufLen]
	b.word.buf = utf8.AppendRune(b.word.buf, '.')
	b.word.buf = utf8.AppendRune(b.word.buf, '.')
	b.word.buf = utf8.AppendRune(b.word.buf, '.')
	b.word.lastRune = '.'
	if b.safe.init {
		b.word.width = b.safe.word.advancedWidth
		b.word.headKerning = b.safe.word.headKerning
	} else {
		b.word.width = b.ellipsisAdvance
		b.word.headKerning = 0
	}
	b.flushWord()

	return true
}
