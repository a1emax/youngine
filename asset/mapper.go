package asset

import (
	"sync"

	"github.com/a1emax/youngine/fault"
)

// Mapper maps kinds to providers.
//
// Mapper can be used concurrently.
type Mapper interface {
	Resolver

	// Map associates given kind with given provider.
	Map(kind Kind, provider Provider)

	// Unmap dissociates given kind from provider if any, or does nothing otherwise.
	Unmap(kind Kind)
}

// mapperImpl is the implementation of the [Mapper] interface.
type mapperImpl struct {
	mu            sync.RWMutex
	kindProviders map[Kind]Provider
}

// NewMapper initializes and returns new [Mapper].
func NewMapper() Mapper {
	return &mapperImpl{
		kindProviders: make(map[Kind]Provider),
	}
}

// Resolve implements the [Resolver] interface.
func (m *mapperImpl) Resolve(kind Kind) (Provider, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	provider, ok := m.kindProviders[kind]

	return provider, ok
}

// Map implements the [Mapper] interface.
func (m *mapperImpl) Map(kind Kind, provider Provider) {
	if kind == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if provider == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.kindProviders[kind]; ok {
		panic(fault.Trace(ErrKindRebinding))
	}

	m.kindProviders[kind] = provider
}

// Unmap implements the [Mapper] interface.
func (m *mapperImpl) Unmap(kind Kind) {
	if kind == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.kindProviders, kind)
}
