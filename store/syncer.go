package store

import (
	"context"
	"reflect"
	"sync"

	"github.com/a1emax/youngine/fault"
)

// Syncer syncs data of type T with store.
//
// Syncer can be used concurrently.
type Syncer[T any] interface {

	// Load reads data from store and writes it to associated locker.
	Load(ctx context.Context) error

	// Save reads data from associated locker and writes it to store.
	Save(ctx context.Context) error
}

// syncerImpl is the implementation of the [Syncer] interface.
type syncerImpl[T any] struct {
	mu       sync.Mutex
	locker   Locker[T]
	data     *T
	accessor Accessor[T]
}

// NewSyncer initializes and returns new [Syncer].
func NewSyncer[T any](locker Locker[T], accessor Accessor[T]) Syncer[T] {
	s := &syncerImpl[T]{}

	if !IsCompatibleType(reflect.TypeOf(s.data).Elem()) {
		panic(fault.Trace(ErrIncompatibleType))
	}

	if locker == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if accessor == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	s.locker = locker
	s.data = new(T)
	s.accessor = accessor

	return s
}

// Load implements the [Syncer] interface.
func (s *syncerImpl[T]) Load(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.accessor.Read(ctx, s.data)
	if err != nil {
		return err
	}

	data, unlock := s.locker.Lock()
	*data = *s.data
	unlock()

	return nil
}

// Save implements the [Syncer] interface.
func (s *syncerImpl[T]) Save(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, unlock := s.locker.RLock()
	*s.data = *data
	unlock()

	return s.accessor.Write(ctx, s.data)
}
