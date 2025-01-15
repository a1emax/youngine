package kage

import (
	"context"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/fault"
)

// Shader asset based on Kage language.
type Shader = *ebiten.Shader

// ShaderProvider provides assets of the [Shader] type.
type ShaderProvider interface {
	asset.Provider
}

// NewShaderProvider initializes and returns new [ShaderProvider].
func NewShaderProvider(fetcher asset.Fetcher) ShaderProvider {
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
