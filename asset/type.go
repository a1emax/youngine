package asset

import (
	"context"
	"reflect"

	"github.com/a1emax/youngine/fault"
)

// typeKind is the implementation of the [Kind] interface representing type T.
type typeKind[T any] struct{}

// Kind implements the [Kind] interface.
func (typeKind[T]) Kind() string {
	var zero T
	return reflect.TypeOf(zero).String()
}

// Map associates type T with given provider.
//
// Map can be used concurrently.
func Map[T any](mapper Mapper, provider Provider) {
	if mapper == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	mapper.Map(typeKind[T]{}, provider)
}

// Unmap dissociates type T from provider if any, or does nothing otherwise.
//
// Unmap can be used concurrently.
func Unmap[T any](mapper Mapper) {
	if mapper == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	mapper.Unmap(typeKind[T]{})
}

// Load loads asset of type T with given URI.
//
// Load can be used concurrently.
func Load[T any](ctx context.Context, loader Loader, uri string) (T, UnloadFunc, error) {
	if loader == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	asset, unload, err := loader.Load(ctx, typeKind[T]{}, uri)
	if err != nil {
		var zero T
		return zero, nil, err
	}

	typedAsset, ok := asset.(T)
	if !ok {
		unload()

		panic(fault.Trace(ErrTypeMismatch))
	}

	return typedAsset, unload, nil
}

// MustLoad calls [Load] and panics if error returned.
//
// MustLoad can be used concurrently.
func MustLoad[T any](ctx context.Context, loader Loader, uri string) (T, UnloadFunc) {
	typedAsset, unload, err := Load[T](ctx, loader, uri)
	if err != nil {
		panic(err)
	}

	return typedAsset, unload
}
