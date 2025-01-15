package store

import (
	"sync"

	"github.com/a1emax/youngine/fault"
)

// Locker holds instance of store data of type T for concurrent use.
type Locker[T any] interface {

	// RLock locks data for reading only and returns it and new [UnlockFunc].
	RLock() (*T, UnlockFunc)

	// Lock locks data both for reading and writing and returns it and new [UnlockFunc].
	Lock() (*T, UnlockFunc)
}

// UnlockFunc unlocks store data.
//
// UnlockFunc panics if called more than once.
type UnlockFunc func()

// lockerImpl is the implementation of the [Locker] interface.
type lockerImpl[T any] struct {
	mu   sync.RWMutex
	data *T
}

// NewLocker initializes and returns new [Locker].
func NewLocker[T any]() Locker[T] {
	return &lockerImpl[T]{
		data: new(T),
	}
}

// RLock implements the [Locker] interface.
func (l *lockerImpl[T]) RLock() (*T, UnlockFunc) {
	l.mu.RLock()
	var unlocked bool

	return l.data, func() {
		if unlocked {
			panic(fault.Trace(fault.ErrInvalidUse))
		}
		unlocked = true

		l.mu.RUnlock()
	}
}

// Lock implements the [Locker] interface.
func (l *lockerImpl[T]) Lock() (*T, UnlockFunc) {
	l.mu.Lock()
	var unlocked bool

	return l.data, func() {
		if unlocked {
			panic(fault.Trace(fault.ErrInvalidUse))
		}
		unlocked = true

		l.mu.Unlock()
	}
}
