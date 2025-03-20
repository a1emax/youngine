package textviewer

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/x/textview"
)

// TextViewer extended by trait of type T.
type TextViewer[T any] interface {
	scene.Element[*ebiten.Image, T]
}

// Props associated with [TextViewer].
type Props struct {
	scene.Attrs

	// FontFace to draw text with.
	FontFace font.Face

	// JustifyLines specifies justification of text lines. Default is [textview.JustifyCenter].
	JustifyLines basic.Opt[textview.Justify]

	// JustifyWords specifies justification of text words within line. Default is [textview.JustifyCenter].
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

// textViewerImpl is the implementation of the [TextViewer] interface.
type textViewerImpl[T any] struct {
	scene.BaseElement[*ebiten.Image, T, Props]

	bbox basic.Rect
	text textObj
}

// New initializes and returns new [TextViewer].
func New[T any](traitFunc scene.TraitFunc[T], propsFunc scene.PropsFunc[Props]) TextViewer[T] {
	if traitFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if propsFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	t := &textViewerImpl[T]{}
	t.Init(traitFunc, propsFunc)

	return t
}

// Attrs implements the [scene.Element] interface.
func (t *textViewerImpl[T]) Attrs() scene.Attrs {
	return t.Props().Attrs
}

// Arrange implements the [scene.Element] interface.
func (t *textViewerImpl[T]) Arrange(bbox basic.Rect) {
	t.bbox = bbox

	t.arrangeText()
}

// Draw implements the [scene.Element] interface.
func (t *textViewerImpl[T]) Draw(screen *ebiten.Image) {
	if screen == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	t.drawText(screen)
}

// Dispose implements the [scene.Element] interface.
func (t *textViewerImpl[T]) Dispose() {
	t.disposeText()
}
