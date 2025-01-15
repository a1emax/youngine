package wav

import (
	"bytes"
	"context"

	"github.com/hajimehoshi/ebiten/v2/audio/wav"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/audio"
	"github.com/a1emax/youngine/fault"
)

// Asset is audio in WAV format.
type Asset = *wav.Stream

// Provider provides assets of the [Asset] type.
type Provider interface {
	asset.Provider
}

// NewProvider initializes and returns new [Provider].
//
// If sample rate is positive, source data will be resampled to fit it.
func NewProvider(fetcher asset.Fetcher, sampleRate audio.SampleRate) Provider {
	if fetcher == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return asset.ProviderFunc(func(ctx context.Context, uri string) (any, func(), error) {
		data, err := fetcher.Fetch(ctx, uri)
		if err != nil {
			return nil, nil, err
		}

		var stream *wav.Stream
		if sampleRate > 0 {
			stream, err = wav.DecodeWithSampleRate(int(sampleRate), bytes.NewReader(data))
		} else {
			stream, err = wav.DecodeWithoutResampling(bytes.NewReader(data))
		}
		if err != nil {
			return nil, nil, fault.Trace(err)
		}

		return stream, nil, nil
	})
}
