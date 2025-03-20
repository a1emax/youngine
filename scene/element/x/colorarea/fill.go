package colorarea

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/basic"
)

// fillObj represents fill.
type fillObj struct {
	key   basic.Opt[fillKey]
	image *ebiten.Image
}

// fillKey is used to detect state changes.
type fillKey struct {
	color color.Color
}

// arrangeFill rebuilds fill if it is not built yet or if state has been changed, or does nothing otherwise.
func (c *colorAreaImpl[T]) arrangeFill() {
	key := fillKey{
		color: c.Props().Color,
	}

	if c.fill.key.IsSet() && c.fill.key.Get() == key {
		return
	}

	c.disposeFill()

	if key.color == nil {
		return
	}

	img := ebiten.NewImage(1, 1)
	img.Fill(key.color)

	c.fill.key = basic.SetOpt(key)
	c.fill.image = img
}

// drawFill draws fill on given screen.
func (c *colorAreaImpl[T]) drawFill(screen *ebiten.Image) {
	if !c.fill.key.IsSet() {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(c.bbox.Width(), c.bbox.Height())
	op.GeoM.Translate(c.bbox.Left(), c.bbox.Top())

	screen.DrawImage(c.fill.image, op)
}

// disposeFill disposes fill.
func (c *colorAreaImpl[T]) disposeFill() {
	if !c.fill.key.IsSet() {
		return
	}

	c.fill.image.Deallocate()

	c.fill = fillObj{}
}
