package txt

import (
	"context"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/fault"
)

// Text asset.
type Text = string

// TextProvider provides assets of the [Text] type.
type TextProvider interface {
	asset.Provider
}

// NewTextProvider initializes and returns new [TextProvider].
func NewTextProvider(fetcher asset.Fetcher) TextProvider {
	if fetcher == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return asset.ProviderFunc(func(ctx context.Context, uri string) (any, func(), error) {
		data, err := fetcher.Fetch(ctx, uri)
		if err != nil {
			return nil, nil, err
		}

		return string(data), nil, nil
	})
}
