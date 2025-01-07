package file

import (
	"context"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/store"
)

// Accessor provides access to file containing data of type T.
type Accessor[T any] interface {
	store.Accessor[T]
}

// accessorImpl is the implementation of the [Accessor] interface.
type accessorImpl[T any] struct {
	filePath string
}

// NewAccessor initializes and returns new [Accessor].
func NewAccessor[T any](filePath string) Accessor[T] {
	if filePath == "" {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	return &accessorImpl[T]{
		filePath: filePath,
	}
}

// Read implements the [store.Accessor] interface.
func (a *accessorImpl[T]) Read(ctx context.Context, data *T) error {
	if ctx == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if data == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	buf, err := os.ReadFile(a.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			var zero T
			*data = zero

			return nil
		}

		return fault.Trace(err)
	}

	err = yaml.Unmarshal(buf, data)
	if err != nil {
		return fault.Trace(err)
	}

	return nil
}

// Write implements the [store.Accessor] interface.
func (a *accessorImpl[T]) Write(ctx context.Context, data *T) error {
	if ctx == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if data == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	buf, err := yaml.Marshal(data)
	if err != nil {
		return fault.Trace(err)
	}

	err = os.WriteFile(a.filePath, buf, 0644)
	if err != nil {
		return fault.Trace(err)
	}

	return nil
}
