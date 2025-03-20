package textviewer

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/x/textview"
)

// textObj represents text.
type textObj struct {
	key  basic.Opt[textKey]
	view textview.TextView
}

// textKey is used to detect state changes.
type textKey struct {
	text      string
	fontFace  font.Face
	sizeLimit basic.Vec2
}

// arrangeText rebuilds text if it is not built yet or if state has been changed, or does nothing otherwise.
func (t *textViewerImpl[T]) arrangeText() {
	props := t.Props()

	key := textKey{
		text:      props.Text,
		fontFace:  props.FontFace,
		sizeLimit: t.bbox.Size,
	}

	if t.text.key.IsSet() && t.text.key.Get() == key {
		return
	}

	t.disposeText()

	if key.fontFace == nil || key.text == "" {
		return
	}

	t.text.key = basic.SetOpt(key)
	t.text.view = textview.New(key.text, key.fontFace, key.sizeLimit)
}

// drawText draws text on given screen.
func (t *textViewerImpl[T]) drawText(screen *ebiten.Image) {
	if !t.text.key.IsSet() {
		return
	}

	props := t.Props()

	if props.TextColor == nil {
		return
	}

	dst := screen
	if size := t.text.view.Size(); size.X() > t.bbox.Width() || size.Y() > t.bbox.Height() {
		dst = dst.SubImage(t.bbox.Image()).(*ebiten.Image)
	}

	t.text.view.Draw(t.bbox.Size,
		props.JustifyLines.Or(textview.JustifyCenter),
		props.JustifyWords.Or(textview.JustifyCenter),
		func(s string, fontFace font.Face, x, y basic.Float) {
			x += t.bbox.Left()
			y += t.bbox.Top()
			text.Draw(dst, s, fontFace, int(math.Floor(x)), int(math.Floor(y)), props.TextColor)
		},
	)
}

// disposeText disposes text.
func (t *textViewerImpl[T]) disposeText() {
	t.text = textObj{}
}
