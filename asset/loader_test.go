package asset

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
)

func TestLoader(t *testing.T) {
	testCtx := context.Background()
	testErr := errors.New("test error")

	var countersMu sync.Mutex
	counters := make(map[string]int)
	incCounter := func(uri string) {
		countersMu.Lock()
		defer countersMu.Unlock()

		counters[uri] = counters[uri] + 1
	}
	decCounter := func(uri string) {
		countersMu.Lock()
		defer countersMu.Unlock()

		counters[uri] = counters[uri] - 1
	}

	kind1 := testKind("kind1")
	kind2 := testKind("kind2")
	kind3 := testKind("kind3")

	type type1 string
	type type2 string

	kindProviders := map[Kind]Provider{
		kind1: ProviderFunc(func(ctx context.Context, uri string) (any, func(), error) {
			incCounter(uri)

			return type1(uri), func() {
				decCounter(uri)
			}, nil
		}),
		kind2: ProviderFunc(func(ctx context.Context, uri string) (any, func(), error) {
			incCounter(uri)

			return type2(uri), func() {
				decCounter(uri)
			}, nil
		}),
		kind3: ProviderFunc(func(ctx context.Context, uri string) (any, func(), error) {
			return nil, nil, testErr
		}),
	}

	loader := NewLoader(ResolverFunc(func(kind Kind) (Provider, bool) {
		provider, ok := kindProviders[kind]

		return provider, ok
	}))

	uri1 := "asset1"
	unloadFuncs1 := make([]func(), 30)
	{
		var fail atomic.Bool
		var wg sync.WaitGroup
		for i := range unloadFuncs1 {
			wg.Add(1)
			go func() {
				defer wg.Done()

				asset, unload, err := loader.Load(testCtx, kind1, uri1)
				if err != nil {
					t.Errorf("asset1: unexpected error: %v", err)
					fail.Store(true)

					return
				}

				if asset != type1(uri1) {
					t.Errorf("asset1: unexpected asset: %+v", asset)
					fail.Store(true)

					return
				}

				unloadFuncs1[i] = unload
			}()
		}
		wg.Wait()
		if fail.Load() {
			t.FailNow()
		}
	}

	uri2 := "asset2"
	unloadFuncs2 := make([]func(), 20)
	{
		var fail atomic.Bool
		var wg sync.WaitGroup
		for i := range unloadFuncs2 {
			wg.Add(1)
			go func() {
				defer wg.Done()

				asset, unload, err := loader.Load(testCtx, kind2, uri2)
				if err != nil {
					t.Errorf("asset2: unexpected error: %v", err)
					fail.Store(true)

					return
				}

				if asset != type2(uri2) {
					t.Errorf("asset2: unexpected asset: %+v", asset)
					fail.Store(true)

					return
				}

				unloadFuncs2[i] = unload
			}()
		}
		wg.Wait()
		if fail.Load() {
			t.FailNow()
		}
	}

	uri3 := "asset3"
	{
		var fail atomic.Bool
		var wg sync.WaitGroup
		for range 10 {
			wg.Add(1)
			go func() {
				defer wg.Done()

				_, _, err := loader.Load(testCtx, kind3, uri3)
				if err == nil {
					t.Errorf("asset3: no error")
					fail.Store(true)

					return
				}
				if !errors.Is(err, testErr) {
					t.Errorf("asset3: unexpected error: %v", err)
					fail.Store(true)

					return
				}
			}()
		}
		wg.Wait()
		if fail.Load() {
			t.FailNow()
		}
	}

	if counter := counters[uri1]; counter != 1 {
		t.Fatalf("asset1: wrong counter before unloading: %d", counter)
	}
	{
		var wg sync.WaitGroup
		for _, unload := range unloadFuncs1 {
			wg.Add(1)

			go func() {
				defer wg.Done()

				unload()
			}()
		}
		wg.Wait()
	}
	if counter := counters[uri1]; counter != 0 {
		t.Fatalf("asset1: wrong counter after unloading: %d", counter)
	}

	if counter := counters[uri2]; counter != 1 {
		t.Fatalf("asset2: wrong counter before unloading: %d", counter)
	}
	{
		var wg sync.WaitGroup
		for _, unload := range unloadFuncs2 {
			wg.Add(1)

			go func() {
				defer wg.Done()

				unload()
			}()
		}
		wg.Wait()
	}
	if counter := counters[uri2]; counter != 0 {
		t.Fatalf("asset2: wrong counter after unloading: %d", counter)
	}

	if counter := counters[uri3]; counter != 0 {
		t.Fatalf("asset3: wrong counter: %d", counter)
	}
}
