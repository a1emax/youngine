package asset

import (
	"github.com/a1emax/youngine/fault"
)

// Resolver resolves kinds to providers.
type Resolver interface {

	// Resolve returns provider associated with given kind if any, or false otherwise.
	Resolve(kind Kind) (Provider, bool)
}

// ResolverFunc is the functional implementation of the [Resolver] interface.
type ResolverFunc func(kind Kind) (Provider, bool)

// Resolve implements the [Resolver] interface.
func (f ResolverFunc) Resolve(kind Kind) (Provider, bool) {
	if f == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return f(kind)
}
