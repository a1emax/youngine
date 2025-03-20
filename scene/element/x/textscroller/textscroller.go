package textscroller

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/x/textview"
)

// TextScroller extended by trait of type T.
type TextScroller[T any] interface {
	scene.Element[*ebiten.Image, T]
}

// Config configures [TextScroller].
type Config struct {
	Clock clock.Clock
	Input input.Input
}

// Props associated with [TextScroller].
type Props struct {
	scene.Attrs

	// FontFace to draw text with.
	FontFace font.Face

	// JustifyWords specifies justification of text words within line. Default is [textview.JustifyStart].
	JustifyWords basic.Opt[textview.Justify]

	// Text to display.
	Text string

	// TextColor to draw text with.
	TextColor color.Color
}

// Func returns these props.
func (p Props) Func(Props) Props {
	return p
}

// textScrollerImpl if the implementation of the [TextScroller] interface.
type textScrollerImpl[T any] struct {
	scene.BaseElement[*ebiten.Image, T, Props]

	bbox       basic.Rect
	controller input.Controller[basic.None]
	text       textObj
	scroll     scrollObj
}

// New initializes and returns new [TextScroller].
func New[T any](traitFunc scene.TraitFunc[T], config Config, propsFunc scene.PropsFunc[Props]) TextScroller[T] {
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

	t := &textScrollerImpl[T]{}
	t.Init(traitFunc, propsFunc)

	t.initController(config)

	return t
}

// Attrs implements the [scene.Element] interface.
func (t *textScrollerImpl[T]) Attrs() scene.Attrs {
	return t.Props().Attrs
}

// Prepare implements the [scene.Element] interface.
func (t *textScrollerImpl[T]) Prepare() {
	t.prepareScroll()
}

// Exclude implements the [scene.Element] interface.
func (t *textScrollerImpl[T]) Exclude() {
	t.excludeScroll()
}

// Arrange implements the [scene.Element] interface.
func (t *textScrollerImpl[T]) Arrange(bbox basic.Rect) {
	t.bbox = bbox

	t.arrangeText()
	t.arrangeScroll()
}

// Actuate implements the [scene.Element] interface.
func (t *textScrollerImpl[T]) Actuate() {
	t.controller.Actuate(basic.None{})
}

// Inhibit implements the [scene.Element] interface.
func (t *textScrollerImpl[T]) Inhibit() {
	t.controller.Inhibit()
}

// Update implements the [scene.Element] interface.
func (t *textScrollerImpl[T]) Update() {
	t.updateScroll()
}

// Draw implements the [scene.Element] interface.
func (t *textScrollerImpl[T]) Draw(screen *ebiten.Image) {
	if screen == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	t.drawText(screen)
}

// Dispose implements the [scene.Element] interface.
func (t *textScrollerImpl[T]) Dispose() {
	t.disposeText()
}
