package sfnt

import (
	"context"
	"fmt"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"gopkg.in/yaml.v3"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/x/scope"
)

// FaceAsset is font face associated with [Asset].
type FaceAsset = font.Face

// FaceProvider provides assets of the [FaceAsset] type.
type FaceProvider interface {
	asset.Provider
}

// NewFaceProvider initializes and returns new [FaceProvider].
func NewFaceProvider(fetcher asset.Fetcher, loader asset.Loader) FaceProvider {
	if fetcher == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if loader == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return asset.ProviderFunc(func(ctx context.Context, uri string) (any, func(), error) {
		var fontFace FaceAsset

		teardown, err := scope.Setup(func(lc scope.Lifecycle) error {
			data, err := fetcher.Fetch(ctx, uri)
			if err != nil {
				return err
			}

			var config struct {
				Font    string      `yaml:"font"`
				Size    basic.Float `yaml:"size"`
				DPI     basic.Float `yaml:"dpi"`
				Hitting string      `yaml:"hitting"`
			}
			err = yaml.Unmarshal(data, &config)
			if err != nil {
				return fault.Trace(err)
			}

			var hitting font.Hinting
			switch config.Hitting {
			default:
				return fault.Trace(fmt.Errorf("invalid font face hitting: %s", config.Hitting))
			case "", "none":
				hitting = font.HintingNone
			case "vertical":
				hitting = font.HintingVertical
			case "full":
				hitting = font.HintingFull
			}

			_font, unloadFont, err := asset.Load[Asset](ctx, loader, config.Font)
			if err != nil {
				return err
			}
			lc.Defer(unloadFont)

			fontFace, err = opentype.NewFace(_font, &opentype.FaceOptions{
				Size:    config.Size,
				DPI:     config.DPI,
				Hinting: hitting,
			})
			if err != nil {
				return fault.Trace(err)
			}
			// The opentype.Face.Close method does nothing.

			return nil
		})
		if err != nil {
			return nil, nil, err
		}

		return fontFace, teardown, nil
	})
}
