package colorarea

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// ColorArea extended by trait of type T.
type ColorArea[T any] interface {
	scene.Element[*ebiten.Image, T]
}

// Props associated with [ColorArea].
type Props struct {
	scene.Attrs

	// Color to fill area with.
	Color color.Color
}

// Func returns these props.
func (p Props) Func(Props) Props {
	return p
}

type colorAreaImpl[T any] struct {
	scene.BaseElement[*ebiten.Image, T, Props]

	bbox basic.Rect
	fill fillObj
}

// New initializes and returns new [ColorArea].
func New[T any](traitFunc scene.TraitFunc[T], propsFunc scene.PropsFunc[Props]) ColorArea[T] {
	if traitFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if propsFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	c := &colorAreaImpl[T]{}
	c.Init(traitFunc, propsFunc)

	return c
}

// Attrs implements the [scene.Element] interface.
func (c *colorAreaImpl[T]) Attrs() scene.Attrs {
	return c.Props().Attrs
}

// Arrange implements the [scene.Element] interface.
func (c *colorAreaImpl[T]) Arrange(bbox basic.Rect) {
	c.bbox = bbox

	c.arrangeFill()
}

// Draw implements the [scene.Element] interface.
func (c *colorAreaImpl[T]) Draw(screen *ebiten.Image) {
	if screen == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	c.drawFill(screen)
}

// Dispose implements the [scene.Element] interface.
func (c *colorAreaImpl[T]) Dispose() {
	c.disposeFill()
}
