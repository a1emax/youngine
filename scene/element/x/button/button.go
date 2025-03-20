package button

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/scene"
)

// Button extended by trait of type T.
type Button[T any] interface {
	scene.Element[*ebiten.Image, T]
}

// Config configures [Button].
type Config struct {
	Clock clock.Clock
	Input input.Input
}

// Props associates with [Button].
type Props struct {
	scene.Attrs

	// CornerRadius specifies radius of shape corners. Default is half height.
	CornerRadius basic.Opt[basic.Float]

	// FontFace to draw text with.
	FontFace font.Face

	// PressedColor specifies color to draw shape with when button is pressed.
	PressedColor color.Color

	// PressedTextColor specifies color to draw text with when button is pressed.
	PressedTextColor color.Color

	// PrimaryColor specifies color to draw shape with when button is not pressed.
	PrimaryColor color.Color

	// PrimaryTextColor specifies color to draw text with when button is not pressed.
	PrimaryTextColor color.Color

	// Text to display.
	Text string

	// OnClick, if specified, is called on [ClickEvent].
	OnClick func(event ClickEvent)

	// OnPress, if specified, is called on [PressEvent].
	OnPress func(event PressEvent)
}

// Func returns these props.
func (p Props) Func(Props) Props {
	return p
}

// buttonImpl is the implementation of the [Button] interface.
type buttonImpl[T any] struct {
	scene.BaseElement[*ebiten.Image, T, Props]

	bbox       basic.Rect
	controller input.Controller[basic.None]
	isPressed  bool
	shape      shapeObj
	text       textObj
}

// New initializes and returns new [Button].
func New[T any](traitFunc scene.TraitFunc[T], config Config, propsFunc scene.PropsFunc[Props]) Button[T] {
	if traitFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if config.Input == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if config.Clock == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if propsFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	b := &buttonImpl[T]{}
	b.Init(traitFunc, propsFunc)

	b.initController(config)

	return b
}

// Attrs implements the [scene.Element] interface.
func (b *buttonImpl[T]) Attrs() scene.Attrs {
	return b.Props().Attrs
}

// Arrange implements the [scene.Element] interface.
func (b *buttonImpl[T]) Arrange(bbox basic.Rect) {
	b.bbox = bbox

	b.arrangeShape()
	b.arrangeText()
}

// Actuate implements the [scene.Element] interface.
func (b *buttonImpl[T]) Actuate() {
	b.controller.Actuate(basic.None{})
}

// Inhibit implements the [scene.Element] interface.
func (b *buttonImpl[T]) Inhibit() {
	b.controller.Inhibit()
}

// Draw implements the [scene.Element] interface.
func (b *buttonImpl[T]) Draw(screen *ebiten.Image) {
	if screen == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	b.drawShape(screen)
	b.drawText(screen)
}

// Dispose implements the [scene.Element] interface.
func (b *buttonImpl[T]) Dispose() {
	b.disposeShape()
	b.disposeText()
}
