package textview

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"golang.org/x/image/font"
)

// TextView represents drawable text.
type TextView struct {
	fontFace font.Face
	lines    []textViewLine
	width    Fixed
	height   Fixed
}

// textViewLine represents line.
type textViewLine struct {
	str           string
	words         []textViewWord
	width         Fixed
	unspacedWidth Fixed
}

// textViewWord represents word within line.
type textViewWord struct {
	low, high int
	width     Fixed
}

// New initializes and returns new [TextView] based on given font face.
//
// Given size limit is soft:
//   - its x component limits width of lines, but each non-empty line will contain at least one rune;
//   - its y component limits height of entire view, but view will contain at least one line
//     (use zero for single line or [basic.PosInf] for unlimited height).
//
// When overflowing size limit, text will be clipped and extended by ellipsis.
func New(text string, fontFace font.Face, sizeLimit basic.Vec2) TextView {
	if fontFace == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	b := builder{
		view: TextView{
			fontFace: fontFace,
		},
		text:      text,
		sizeLimit: sizeLimit,
	}

	b.build()

	return b.view
}

// IsNil reports whether view is nil.
func (v TextView) IsNil() bool {
	return v.fontFace == nil
}

// Size returns size of view.
func (v TextView) Size() basic.Vec2 {
	return basic.Vec2{
		FixedToFloat(v.width),
		FixedToFloat(v.height),
	}
}
