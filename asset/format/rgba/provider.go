package rgba

import (
	"context"

	"gopkg.in/yaml.v3"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/x/colors"
)

// Asset is color in RGBA format.
type Asset = colors.RGBA

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

		var config struct {
			Color  uint32       `yaml:"color"`
			Alpha  *uint8       `yaml:"alpha"`
			FAlpha *basic.Float `yaml:"falpha"`
			Premul bool         `yaml:"premul"`
		}
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			return nil, nil, fault.Trace(err)
		}

		color := colors.RGBA{
			uint8((config.Color >> 16) & 0xFF),
			uint8((config.Color >> 8) & 0xFF),
			uint8(config.Color & 0xFF),
			0xFF,
		}

		if config.Alpha != nil {
			color[3] = *config.Alpha
		} else if config.FAlpha != nil {
			color[3] = uint8(0xFF * *config.FAlpha)
		}

		if config.Premul {
			color = color.Premul()
		}

		return color, nil, nil
	})
}
