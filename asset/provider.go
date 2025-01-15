package asset

import (
	"context"

	"github.com/a1emax/youngine/fault"
)

// Provider provides assets.
type Provider interface {

	// Provide returns asset with given URI. Second output parameter, if is not nil, is function that disposes asset.
	Provide(ctx context.Context, uri string) (any, func(), error)
}

// ProviderFunc is the functional implementation of the [Provider] interface.
type ProviderFunc func(ctx context.Context, uri string) (any, func(), error)

// Provide implements the [Provider] interface.
func (f ProviderFunc) Provide(ctx context.Context, uri string) (any, func(), error) {
	if f == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return f(ctx, uri)
}
