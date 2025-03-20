package button

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/x/roundrect"
)

// shapeObj represents shape.
type shapeObj struct {
	key   basic.Opt[shapeKey]
	image *ebiten.Image
}

// shapeKey is used to detect state changes.
type shapeKey struct {
	size         basic.Vec2
	cornerRadius basic.Float
}

// arrangeShape rebuilds shape if it is not built yet or if state has been changed, or does nothing otherwise.
func (b *buttonImpl[T]) arrangeShape() {
	var cornerRadius basic.Float
	if v := b.Props().CornerRadius; v.IsSet() {
		cornerRadius = v.Get()
	} else {
		cornerRadius = b.bbox.Height() / 2
	}

	key := shapeKey{
		size:         b.bbox.Size,
		cornerRadius: cornerRadius,
	}

	if b.shape.key.IsSet() && b.shape.key.Get() == key {
		return
	}

	b.disposeShape()

	bmp := roundrect.Fill(key.size.X(), key.size.Y(), key.cornerRadius, key.cornerRadius)
	img := ebiten.NewImage(bmp.Width(), bmp.Height())
	img.WritePixels(bmp.Data())

	b.shape.key = basic.SetOpt(key)
	b.shape.image = img
}

// drawShape draws shape on given screen.
func (b *buttonImpl[T]) drawShape(screen *ebiten.Image) {
	if !b.shape.key.IsSet() {
		return
	}

	props := b.Props()

	clr := b.color(props.PrimaryColor, props.PressedColor)
	if clr == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.bbox.Left(), b.bbox.Top())
	op.ColorScale.ScaleWithColor(clr)

	screen.DrawImage(b.shape.image, op)
}

// disposeShape disposes shape.
func (b *buttonImpl[T]) disposeShape() {
	if !b.shape.key.IsSet() {
		return
	}

	b.shape.image.Deallocate()

	b.shape = shapeObj{}
}
