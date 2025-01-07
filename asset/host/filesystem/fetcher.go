package filesystem

import (
	"context"
	"io/fs"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/fault"
)

// Fetcher based on file system.
type Fetcher interface {
	asset.Fetcher
}

// NewFetcher initializes and returns new [Fetcher].
func NewFetcher(fileSystem fs.FS) Fetcher {
	if fileSystem == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return asset.FetcherFunc(func(ctx context.Context, path string) ([]byte, error) {
		data, err := fs.ReadFile(fileSystem, path)
		if err != nil {
			return nil, fault.Trace(err)
		}

		return data, nil
	})
}
