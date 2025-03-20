package asset

import (
	"context"
	"sync"

	"github.com/a1emax/youngine/fault"
)

// Loader loads assets.
//
// Loader can be used concurrently.
type Loader interface {

	// Load receives asset with given URI from provider associated with given kind and puts it to cache,
	// if it has not been received and cached yet, and then increases its reference counter by one and
	// returns asset and new [UnloadFunc].
	Load(ctx context.Context, kind Kind, uri string) (any, UnloadFunc, error)
}

// UnloadFunc decreases reference counter of associated asset by one and removes it from cache, disposing it
// if necessary, if there are no more references.
//
// UnloadFunc panics if called more than once.
type UnloadFunc func()

// loaderImpl is the implementation of the [Loader] interface.
type loaderImpl struct {
	resolver Resolver
	mu       sync.Mutex
	entries  map[string]*loaderEntry
}

// loaderEntry is internal representation of loaded asset.
type loaderEntry struct {
	kind    Kind // does not change
	rc      int  // protected by loaderImpl.mu
	ready   chan struct{}
	asset   any
	dispose func()
	err     error
}

// NewLoader initializes and returns new [Loader].
func NewLoader(resolver Resolver) Loader {
	if resolver == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &loaderImpl{
		resolver: resolver,
		entries:  make(map[string]*loaderEntry),
	}
}

// Load implements the [Loader] interface.
func (l *loaderImpl) Load(ctx context.Context, kind Kind, uri string) (any, UnloadFunc, error) {
	if kind == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	entry, receive, err := l.entry(kind, uri)
	if err != nil {
		return nil, nil, err
	}

	if receive {
		defer func() {
			close(entry.ready)

			if entry.err != nil {
				l.mu.Lock()
				delete(l.entries, uri)
				l.mu.Unlock()
			}
		}()

		entry.err = ErrNotLoaded

		provider, ok := l.resolver.Resolve(kind)
		if !ok {
			entry.err = ErrUnknownKind

			return nil, nil, entry.err
		}

		asset, dispose, err := provider.Provide(ctx, uri)
		if err != nil {
			entry.err = err

			return nil, nil, entry.err
		}

		entry.asset = asset
		entry.dispose = dispose
		entry.err = nil
	} else {
		<-entry.ready

		if entry.err != nil {
			return nil, nil, entry.err
		}
	}

	return entry.asset, l.newUnloadFunc(uri, entry), nil
}

// entry returns new entry for receiving asset of given kind with given URI or existing one for waiting for it.
func (l *loaderImpl) entry(kind Kind, uri string) (entry *loaderEntry, receive bool, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	entry, ok := l.entries[uri]
	if ok {
		if kind != entry.kind {
			return nil, false, ErrKindMismatch
		}

		entry.rc++

		return entry, false, nil
	}

	entry = &loaderEntry{
		kind:  kind,
		rc:    1,
		ready: make(chan struct{}),
	}
	l.entries[uri] = entry

	return entry, true, nil
}

// newUnloadFunc returns new [UnloadFunc] for asset with given URI associated with given entry.
func (l *loaderImpl) newUnloadFunc(uri string, entry *loaderEntry) UnloadFunc {
	return func() {
		if entry == nil {
			panic(fault.Trace(fault.ErrInvalidUse))
		}
		defer func() {
			entry = nil
		}()

		live := true
		l.mu.Lock()
		if entry.rc--; entry.rc == 0 {
			delete(l.entries, uri)
			live = false
		}
		l.mu.Unlock()

		if live {
			return
		}

		if entry.dispose != nil {
			entry.dispose()
		}
	}
}
