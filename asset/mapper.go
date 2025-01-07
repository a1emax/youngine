package asset

import (
	"github.com/a1emax/youngine/fault"
)

// Mapper maps kinds to providers.
type Mapper interface {
	Binder
	Resolver
}

// mapperImpl is the implementation of the [Mapper] interface.
type mapperImpl struct {
	kindProviders map[Kind]Provider
}

// NewMapper initializes and returns new [Mapper].
func NewMapper() Mapper {
	return &mapperImpl{
		kindProviders: make(map[Kind]Provider),
	}
}

// Bind implements the [Binder] interface.
func (m *mapperImpl) Bind(kind Kind, provider Provider) {
	if kind == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if provider == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	if _, ok := m.kindProviders[kind]; ok {
		panic(fault.Trace(ErrKindRebinding))
	}

	m.kindProviders[kind] = provider
}

// Resolve implements the [Resolver] interface.
func (m *mapperImpl) Resolve(kind Kind) Provider {
	provider, ok := m.kindProviders[kind]
	if !ok {
		panic(fault.Trace(ErrUnknownKind))
	}

	return provider
}
