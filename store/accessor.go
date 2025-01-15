package store

import (
	"context"
)

// Accessor provides access to store of data of type T.
type Accessor[T any] interface {

	// Read reads data from store.
	Read(ctx context.Context, data *T) error

	// Write writes data to store.
	Write(ctx context.Context, data *T) error
}
