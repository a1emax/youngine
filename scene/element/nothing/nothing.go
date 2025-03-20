package nothing

import (
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
)

// Nothing displayed on screen of type S and extended by trait of type T.
type Nothing[S, T any] interface {
	scene.Element[S, T]
}

// Props associated with [Nothing].
type Props struct {
	scene.Attrs
}

// Func returns these props.
func (p Props) Func(Props) Props {
	return p
}

// nothingImpl is the implementation of the [Nothing] interface.
type nothingImpl[S, T any] struct {
	scene.BaseElement[S, T, Props]
}

// New initializes and returns new [Nothing].
func New[S, T any](traitFunc scene.TraitFunc[T], propsFunc scene.PropsFunc[Props]) Nothing[S, T] {
	if traitFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if propsFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	n := &nothingImpl[S, T]{}
	n.Init(traitFunc, propsFunc)

	return n
}

// Attrs implements the [scene.Element] interface.
func (n *nothingImpl[S, T]) Attrs() scene.Attrs {
	return n.Props().Attrs
}
