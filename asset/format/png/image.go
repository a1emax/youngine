package png

import (
	"bytes"
	"context"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/fault"
)

// Image asset based on PNG format.
type Image = *ebiten.Image

// ImageProvider provides assets of the [Image] type.
type ImageProvider interface {
	asset.Provider
}

// NewImageProvider initializes and returns new [ImageProvider].
func NewImageProvider(fetcher asset.Fetcher) ImageProvider {
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
