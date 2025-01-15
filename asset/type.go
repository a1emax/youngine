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

// Bind associates type T with given provider using given binder.
func Bind[T any](binder Binder, provider Provider) {
	if binder == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	binder.Bind(typeKind[T]{}, provider)
}

// Load loads asset of type T with given URI using given loader.
func Load[T any](ctx context.Context, loader Loader, uri string) (T, UnloadFunc, error) {
	if loader == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	untypedAsset, unload, err := loader.Load(ctx, typeKind[T]{}, uri)
	if err != nil {
		var zero T
		return zero, nil, err
	}

	typedAsset, ok := untypedAsset.(T)
	if !ok {
		unload()

		panic(fault.Trace(ErrTypeMismatch))
	}

	return typedAsset, unload, nil
}

// MustLoad calls [Load] and panics if error returned.
func MustLoad[T any](ctx context.Context, loader Loader, uri string) (T, UnloadFunc) {
	typedAsset, unload, err := Load[T](ctx, loader, uri)
	if err != nil {
		panic(err)
	}

	return typedAsset, unload
}
