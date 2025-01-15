package ringbuffer

import (
	"github.com/a1emax/youngine/fault"
)

// RingBuffer is ring buffer of values of type T.
type RingBuffer[T any] struct {
	*ringBufferInst[T]
}

// ReadOnly is [RingBuffer] with read-only access.
type ReadOnly[T any] struct {
	*ringBufferInst[T]
}

// ringBufferInst is the internal state of the [RingBuffer] type (shared between copies).
//
// ringBufferInst methods do not change it.
type ringBufferInst[T any] struct {
	array  []T // Its length is buffer capacity.
	offset int // Index of the first element in array.
	len    int
}

// New initializes and returns new [RingBuffer] of given capacity.
func New[T any](cap int) RingBuffer[T] {
	return RingBuffer[T]{
		&ringBufferInst[T]{
			array:  make([]T, cap),
			offset: 0,
			len:    0,
		},
	}
}

// ReadOnly returns buffer with read-only access.
func (b RingBuffer[T]) ReadOnly() ReadOnly[T] {
	return ReadOnly[T](b)
}

// Copy returns copy of buffer.
func (b *ringBufferInst[T]) Copy() RingBuffer[T] {
	if b.IsNil() {
		return RingBuffer[T]{}
	}

	result := New[T](len(b.array))
	copy(result.array, b.array)
	result.offset = b.offset
	result.len = b.len

	return result
}

// IsNil reports whether buffer is nil.
func (b *ringBufferInst[T]) IsNil() bool {
	return b == nil
}

// Cap returns capacity of buffer.
func (b *ringBufferInst[T]) Cap() int {
	if b.IsNil() {
		return 0
	}

	return len(b.array)
}

// Len returns number of elements.
func (b *ringBufferInst[T]) Len() int {
	if b.IsNil() {
		return 0
	}

	return b.len
}

// Get returns value of element with given index.
func (b *ringBufferInst[T]) Get(index int) T {
	if b.IsNil() || index < 0 || index >= b.len {
		panic(fault.Trace(fault.ErrIndexOutOfRange))
	}

	index += b.offset
	if index >= len(b.array) {
		index -= len(b.array)
	}

	return b.array[index]
}

// Append appends one element with given value to buffer's end. If resulting number of elements exceeds capacity,
// the first element will be removed from buffer.
func (b RingBuffer[T]) Append(value T) {
	if b.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	var index int
	if b.len < len(b.array) {
		index = b.offset + b.len
		if index >= len(b.array) {
			index -= len(b.array)
		}

		b.len++
	} else { // b.len == len(b.array)
		index = b.offset

		b.offset++
		if b.offset == len(b.array) {
			b.offset = 0
		}
	}

	b.array[index] = value
}
