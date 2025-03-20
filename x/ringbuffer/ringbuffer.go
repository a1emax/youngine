package ringbuffer

import (
	"iter"

	"github.com/a1emax/youngine/fault"
)

// RingBuffer of elements with values of type T.
//
// Values of this type are references to shared internal state. Use the Copy method to make separate copy.
type RingBuffer[T any] struct {
	*ringBufferInst[T]
}

// ringBufferInst is the internal state of the [RingBuffer] type.
type ringBufferInst[T any] struct {
	array  []T // Its length is buffer cap.
	offset int // Index of the first element in array.
	len    int
}

// New initializes and returns new [RingBuffer] of given capacity.
func New[T any](cap int) RingBuffer[T] {
	if cap < 0 {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	return RingBuffer[T]{
		&ringBufferInst[T]{
			array:  make([]T, cap),
			offset: 0,
			len:    0,
		},
	}
}

// IsNil reports whether buffer is nil.
func (b RingBuffer[T]) IsNil() bool {
	return b.ringBufferInst == nil
}

// Copy returns copy of buffer.
func (b RingBuffer[T]) Copy() RingBuffer[T] {
	if b.IsNil() {
		return RingBuffer[T]{}
	}

	result := RingBuffer[T]{
		&ringBufferInst[T]{
			array:  make([]T, len(b.array)),
			offset: b.offset,
			len:    b.len,
		},
	}

	copy(result.array, b.array)

	return result
}

// All returns iterator that iterates over all elements in direct order, producing their values.
func (b RingBuffer[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		if b.IsNil() {
			return
		}

		for i := 0; i < b.len; i++ {
			j := i + b.offset
			if j >= len(b.array) {
				j -= len(b.array)
			}

			if !yield(b.array[j]) {
				break
			}
		}
	}
}

// Indexed returns iterator that iterates over all elements in direct order, producing both their indexes and values.
func (b RingBuffer[T]) Indexed() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		if b.IsNil() {
			return
		}

		for i := 0; i < b.len; i++ {
			j := i + b.offset
			if j >= len(b.array) {
				j -= len(b.array)
			}

			if !yield(i, b.array[j]) {
				break
			}
		}
	}
}

// Backward returns iterator that iterates over all elements in backward order, producing their values.
func (b RingBuffer[T]) Backward() iter.Seq[T] {
	return func(yield func(T) bool) {
		if b.IsNil() {
			return
		}

		for i := b.len - 1; i >= 0; i-- {
			j := i + b.offset
			if j >= len(b.array) {
				j -= len(b.array)
			}

			if !yield(b.array[j]) {
				break
			}
		}
	}
}

// Cap returns capacity of buffer.
func (b RingBuffer[T]) Cap() int {
	if b.IsNil() {
		return 0
	}

	return len(b.array)
}

// Len returns number of elements.
func (b RingBuffer[T]) Len() int {
	if b.IsNil() {
		return 0
	}

	return b.len
}

// Get returns value of element with given index.
func (b RingBuffer[T]) Get(index int) T {
	if b.IsNil() || index < 0 || index >= b.len {
		panic(fault.Trace(fault.ErrIndexOutOfRange))
	}

	index += b.offset
	if index >= len(b.array) {
		index -= len(b.array)
	}

	return b.array[index]
}

// Set sets given value for element with given index.
func (b RingBuffer[T]) Set(index int, value T) {
	if b.IsNil() || index < 0 || index >= b.len {
		panic(fault.Trace(fault.ErrIndexOutOfRange))
	}

	index += b.offset
	if index >= len(b.array) {
		index -= len(b.array)
	}

	b.array[index] = value
}

// Append appends element with given value to buffer's end. If resulting number of elements exceeds capacity,
// the first element will be removed from buffer.
func (b RingBuffer[T]) Append(value T) {
	if b.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	var j int
	if b.len < len(b.array) {
		j = b.offset + b.len
		if j >= len(b.array) {
			j -= len(b.array)
		}

		b.len++
	} else { // b.len == len(b.array)
		j = b.offset

		b.offset++
		if b.offset == len(b.array) {
			b.offset = 0
		}
	}

	b.array[j] = value
}
