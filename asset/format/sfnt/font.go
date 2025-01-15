package sfnt

import (
	"context"

	"golang.org/x/image/font/sfnt"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/fault"
)

// Font asset based on SFNT format.
type Font = *sfnt.Font

// FontProvider provides assets of the [Font] type.
type FontProvider interface {
	asset.Provider
}

// NewFontProvider initializes and returns new [FontProvider].
func NewFontProvider(fetcher asset.Fetcher) FontProvider {
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
