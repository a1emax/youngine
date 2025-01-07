package asset

// Resolver resolves kinds to providers.
type Resolver interface {

	// Resolve returns provider associated with given kind.
	Resolve(kind Kind) Provider
}
