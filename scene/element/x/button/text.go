package button

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
	text       string
	fontFace   font.Face
	widthLimit basic.Float
}

// arrangeText rebuilds text if it is not built yet or if state has been changed, or does nothing otherwise.
func (b *buttonImpl[T]) arrangeText() {
	props := b.Props()

	key := textKey{
		text:       props.Text,
		fontFace:   props.FontFace,
		widthLimit: b.bbox.Width(),
	}

	if b.text.key.IsSet() && b.text.key.Get() == key {
		return
	}

	b.disposeText()

	if key.fontFace == nil || key.text == "" {
		return
	}

	b.text.key = basic.SetOpt(key)
	b.text.view = textview.New(key.text, key.fontFace, basic.Vec2{key.widthLimit, 0.0})
}

// drawText draws text on given screen.
func (b *buttonImpl[T]) drawText(screen *ebiten.Image) {
	if !b.text.key.IsSet() {
		return
	}

	props := b.Props()

	clr := b.color(props.PrimaryTextColor, props.PressedTextColor)
	if clr == nil {
		return
	}

	dst := screen
	if size := b.text.view.Size(); size.X() > b.bbox.Width() || size.Y() > b.bbox.Height() {
		dst = dst.SubImage(b.bbox.Image()).(*ebiten.Image)
	}

	b.text.view.Draw(b.bbox.Size, textview.JustifyCenter, textview.JustifyCenter,
		func(s string, fontFace font.Face, x, y basic.Float) {
			x += b.bbox.Left()
			y += b.bbox.Top()
			text.Draw(dst, s, fontFace, int(math.Floor(x)), int(math.Floor(y)), clr)
		},
	)
}

// disposeText disposes text.
func (b *buttonImpl[T]) disposeText() {
	b.text = textObj{}
}
