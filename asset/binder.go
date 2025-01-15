package asset

// Binder binds kinds to providers.
type Binder interface {

	// Bind associates given kind with given provider.
	Bind(kind Kind, provider Provider)
}
