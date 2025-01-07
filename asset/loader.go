package asset

import (
	"context"

	"github.com/a1emax/youngine/fault"
)

// Loader loads assets.
type Loader interface {

	// Load receives asset with given URI from provider associated with given kind and puts it to cache,
	// if it has not been received and cached yet, and then increases number of its loads by one and
	// returns asset and new [UnloadFunc].
	Load(ctx context.Context, kind Kind, uri string) (any, UnloadFunc, error)
}

// UnloadFunc decreases number of loads of associated asset by one and removes it from cache, disposing it
// if necessary, if resulting number is zero.
//
// UnloadFunc panics if called more than once.
type UnloadFunc func()

// loaderImpl is the implementation of the [Loader] interface.
type loaderImpl struct {
	resolver Resolver
	cache    map[loaderCacheKey]*loaderCacheEntry
}

// loaderCacheKey represents asset URI.
type loaderCacheKey = string

// loaderCacheEntry represents asset for internal purposes.
type loaderCacheEntry struct {
	any

	kind      Kind
	loadCount int
	dispose   func()
}

// NewLoader initializes and returns new [Loader].
func NewLoader(resolver Resolver) Loader {
	if resolver == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &loaderImpl{
		resolver: resolver,
		cache:    make(map[string]*loaderCacheEntry),
	}
}

// Load implements the [Loader] interface.
func (l *loaderImpl) Load(ctx context.Context, kind Kind, uri string) (any, UnloadFunc, error) {
	if kind == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	entry, ok := l.cache[uri]
	if ok {
		if kind != entry.kind {
			panic(fault.Trace(ErrKindMismatch))
		}

		entry.loadCount++
	} else {
		untypedAsset, dispose, err := l.resolver.Resolve(kind).Provide(ctx, uri)
		if err != nil {
			return nil, nil, err
		}

		entry = &loaderCacheEntry{untypedAsset, kind, 1, dispose}

		l.cache[uri] = entry
	}

	return entry.any, func() {
		if entry == nil {
			panic(fault.Trace(fault.ErrInvalidUse))
		}
		defer func() {
			entry = nil
		}()

		if entry.loadCount--; entry.loadCount == 0 {
			delete(l.cache, uri)

			if entry.dispose != nil {
				entry.dispose()
			}
		}
	}, nil
}
