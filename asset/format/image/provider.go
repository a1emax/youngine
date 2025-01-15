package image

import (
	"bytes"
	"context"
	_ "image/jpeg"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/fault"
)

// Asset is image in one of registered formats. JPEG and PNG formats are registered by default.
type Asset = *ebiten.Image

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

		image, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(data))
		if err != nil {
			return nil, nil, fault.Trace(err)
		}

		return image, image.Deallocate, nil
	})
}
