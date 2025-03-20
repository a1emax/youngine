package list

import (
	"iter"

	"github.com/a1emax/youngine/fault"
)

// List is doubly linked list of elements with values of type T.
//
// Values of this type are references to shared internal state. Use the Copy method to make separate copy.
type List[T any] struct {
	*listInst[T]
}

// listInst is the internal state of the [List] type.
type listInst[T any] struct {
	first, last *listEntry[T]
	len         int
}

// listEntry is internal representation of list element.
type listEntry[T any] struct {

	// list reference is placed here both to avoid dynamic allocations when returning markers
	// and to make it possible to invalidate them on deletion. But this solution has extra
	// cost of 8 bytes per entry and requires refactoring if it is possible.
	list List[T]

	prev, next *listEntry[T]
	value      T
}

// Marker of [List] element.
type Marker[T any] struct {
	entry *listEntry[T]
}

// IsNil reports whether marker is nil.
func (m Marker[T]) IsNil() bool {
	return m.entry == nil
}

// New initializes and returns new [List].
func New[T any]() List[T] {
	return List[T]{
		&listInst[T]{},
	}
}

// IsNil reports whether list is nil.
func (l List[T]) IsNil() bool {
	return l.listInst == nil
}

// Copy returns copy of list.
func (l List[T]) Copy() List[T] {
	if l.IsNil() {
		return List[T]{}
	}

	result := New[T]()

	for entry := l.first; entry != nil; entry = entry.next {
		result.putLast(result.newEntry(entry.value))
	}

	result.len = l.len

	return result
}

// All returns iterator that iterates over all elements in direct order, producing their values.
func (l List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		if l.IsNil() {
			return
		}

		for entry := l.first; entry != nil; entry = entry.next {
			if !yield(entry.value) {
				break
			}
		}
	}
}

// Backward returns iterator that iterates over all elements in backward order, producing their values.
func (l List[T]) Backward() iter.Seq[T] {
	return func(yield func(T) bool) {
		if l.IsNil() {
			return
		}

		for entry := l.last; entry != nil; entry = entry.prev {
			if !yield(entry.value) {
				break
			}
		}
	}
}

// Len returns number of elements.
func (l List[T]) Len() int {
	if l.IsNil() {
		return 0
	}

	return l.len
}

// First returns marker of the first element, or nil if list is empty.
func (l List[T]) First() Marker[T] {
	if l.IsNil() {
		return Marker[T]{}
	}

	return Marker[T]{l.first}
}

// Last returns marker of the last element, or nil if the list is empty.
func (l List[T]) Last() Marker[T] {
	if l.IsNil() {
		return Marker[T]{}
	}

	return Marker[T]{l.last}
}

// Prev returns marker of element previous to one with given marker.
func (l List[T]) Prev(marker Marker[T]) Marker[T] {
	if marker.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if marker.entry.list != l {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	if marker.entry.prev == nil {
		return Marker[T]{}
	}

	return Marker[T]{marker.entry.prev}
}

// Next returns marker of element next to one with given marker.
func (l List[T]) Next(marker Marker[T]) Marker[T] {
	if marker.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if marker.entry.list != l {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	if marker.entry.next == nil {
		return Marker[T]{}
	}

	return Marker[T]{marker.entry.next}
}

// Get returns value of element with given marker.
func (l List[T]) Get(marker Marker[T]) T {
	if marker.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if marker.entry.list != l {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	return marker.entry.value
}

// Set sets given value for element with given marker.
func (l List[T]) Set(marker Marker[T], value T) {
	if marker.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if marker.entry.list != l {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	marker.entry.value = value
}

// Prepend inserts new element with given value at front of list, and returns its marker.
func (l List[T]) Prepend(value T) Marker[T] {
	if l.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	entry := l.newEntry(value)
	l.putFirst(entry)

	l.len++

	return Marker[T]{entry}
}

// Append inserts new element with given value at back of list, and returns its marker.
func (l List[T]) Append(value T) Marker[T] {
	if l.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	entry := l.newEntry(value)
	l.putLast(entry)

	l.len++

	return Marker[T]{entry}
}

// InsertBefore inserts new element with given value before pivot one with given marker, and returns its marker.
func (l List[T]) InsertBefore(pivot Marker[T], value T) Marker[T] {
	if l.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if pivot.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if pivot.entry.list != l {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	entry := l.newEntry(value)
	l.putBefore(pivot.entry, entry)

	l.len++

	return Marker[T]{entry}
}

// InsertAfter inserts new element with given value after pivot one with given marker, and returns its marker.
func (l List[T]) InsertAfter(pivot Marker[T], value T) Marker[T] {
	if l.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if pivot.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if pivot.entry.list != l {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	entry := l.newEntry(value)
	l.putAfter(pivot.entry, entry)

	l.len++

	return Marker[T]{entry}
}

// Delete deletes element with given marker from list and returns its value.
func (l List[T]) Delete(marker Marker[T]) T {
	if l.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if marker.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if marker.entry.list != l {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	value := marker.entry.value

	l.cut(marker.entry)
	*marker.entry = listEntry[T]{}

	l.len--

	return value
}

// newEntry initializes and returns new entry.
func (l List[T]) newEntry(value T) *listEntry[T] {
	return &listEntry[T]{
		list:  l,
		value: value,
	}
}

// putFirst puts given new entry at front of list.
func (l List[T]) putFirst(entry *listEntry[T]) {
	if l.first == nil {
		l.first = entry
		l.last = entry
		// entry.prev = nil
		// entry.next = nil
	} else {
		l.putBefore(l.first, entry)
	}
}

// putLast puts given new entry at back of list.
func (l List[T]) putLast(entry *listEntry[T]) {
	if l.last == nil {
		l.putFirst(entry)
	} else {
		l.putAfter(l.last, entry)
	}
}

// putBefore puts given entry before given pivot one.
func (l List[T]) putBefore(pivot, entry *listEntry[T]) {
	entry.next = pivot

	if pivot.prev == nil {
		// entry.prev = nil
		l.first = entry
	} else {
		entry.prev = pivot.prev
		pivot.prev.next = entry
	}

	pivot.prev = entry
}

// putAfter puts given entry after given pivot one.
func (l List[T]) putAfter(pivot, entry *listEntry[T]) {
	entry.prev = pivot

	if pivot.next == nil {
		// entry.next = nil
		l.last = entry
	} else {
		entry.next = pivot.next
		pivot.next.prev = entry
	}

	pivot.next = entry
}

// cut cuts given entry from list.
func (l List[T]) cut(entry *listEntry[T]) {
	if entry.prev == nil {
		l.first = entry.next
	} else {
		entry.prev.next = entry.next
	}

	if entry.next == nil {
		l.last = entry.prev
	} else {
		entry.next.prev = entry.prev
	}

	// entry.prev = nil
	// entry.next = nil
}
