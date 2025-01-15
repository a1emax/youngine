package asset

import (
	"context"

	"github.com/a1emax/youngine/fault"
)

// Fetcher fetches asset data. It can be used to abstract data source from providers.
type Fetcher interface {

	// Fetch returns raw data located at given path.
	Fetch(ctx context.Context, path string) ([]byte, error)
}

// FetcherFunc is the functional implementation of the [Fetcher] interface.
type FetcherFunc func(ctx context.Context, path string) ([]byte, error)

// Fetch implements the [Fetcher] interface.
func (f FetcherFunc) Fetch(ctx context.Context, path string) ([]byte, error) {
	if f == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return f(ctx, path)
}
