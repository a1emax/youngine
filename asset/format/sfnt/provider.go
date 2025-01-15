package sfnt

import (
	"context"

	"golang.org/x/image/font/sfnt"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/fault"
)

// Asset is font in SFNT format.
type Asset = *sfnt.Font

// Provider provides assets of the [Asset] type.
type Provider interface {
	asset.Provider
}

// NewProvider initializes and returns new [Provider].
func NewProvider(fetcher asset.Fetcher) Provider {
	if fetcher == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return asset.ProviderFunc(func(ctx context.Context, uri string) (any, func(), error) {
		data, err := fetcher.Fetch(ctx, uri)
		if err != nil {
			return nil, nil, err
		}

		font, err := sfnt.Parse(data)
		if err != nil {
			return nil, nil, fault.Trace(err)
		}

		return font, nil, nil
	})
}
