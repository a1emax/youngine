package kage

import (
	"context"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/fault"
)

// Asset is shader written in Kage language.
type Asset = *ebiten.Shader

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

		shader, err := ebiten.NewShader(data)
		if err != nil {
			return nil, nil, fault.Trace(err)
		}

		return shader, shader.Deallocate, nil
	})
}
