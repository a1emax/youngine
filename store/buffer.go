package store

import (
	"reflect"

	"github.com/a1emax/youngine/fault"
)

// Buffer holds instance of store data of type T for non-concurrent use.
type Buffer[T any] interface {

	// Data returns data.
	Data() *T

	// Pull reads data from associated locker and writes it to buffer.
	Pull()

	// Push reads data from buffer and writes it to associated locker.
	Push()
}

// bufferImpl is the implementation of the [Buffer] interface.
type bufferImpl[T any] struct {
	locker Locker[T]
	data   *T
}

// NewBuffer initializes and returns new [Buffer].
func NewBuffer[T any](locker Locker[T]) Buffer[T] {
	b := &bufferImpl[T]{}

	if !IsCompatibleType(reflect.TypeOf(b.data).Elem()) {
		panic(fault.Trace(ErrIncompatibleType))
	}

	if locker == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	b.locker = locker
	b.data = new(T)

	return b
}

// Data implements the [Buffer] interface.
func (b *bufferImpl[T]) Data() *T {
	return b.data
}

// Pull implements the [Buffer] interface.
func (b *bufferImpl[T]) Pull() {
	data, unlock := b.locker.RLock()
	*b.data = *data
	unlock()
}

// Push implements the [Buffer] interface.
func (b *bufferImpl[T]) Push() {
	data, unlock := b.locker.Lock()
	*data = *b.data
	unlock()
}
