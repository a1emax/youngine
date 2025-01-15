package list

import (
	"github.com/a1emax/youngine/fault"
)

// List is doubly linked list of values of type T.
type List[T any] struct {
	*listInst[T]
}

// ReadOnly is [List] with read-only access.
type ReadOnly[T any] struct {
	*listInst[T]
}

// listInst is the internal state of the [List] type (shared between copies).
//
// listInst methods do not change it.
type listInst[T any] struct {
	first, last Entry[T]
	len         int
}

// New initializes and returns new [List].
func New[T any]() List[T] {
	return List[T]{
		&listInst[T]{},
	}
}

// ReadOnly returns list with read-only access.
func (l List[T]) ReadOnly() ReadOnly[T] {
	return ReadOnly[T](l)
}

// Copy returns copy of list.
func (l *listInst[T]) Copy() List[T] {
	if l.IsNil() {
		return List[T]{}
	}

	result := New[T]()

	for e := l.first; !e.IsNil(); e = e.next {
		result.putLast(Entry[T]{
			&entryInst[T]{
				list:  result,
				value: e.value,
			},
		})
	}

	result.len = l.len

	return result
}

// IsNil reports whether list is nil.
func (l *listInst[T]) IsNil() bool {
	return l == nil
}

// Len returns number of entries.
func (l *listInst[T]) Len() int {
	if l.IsNil() {
		return 0
	}

	return l.len
}

// First returns the first entry, or nil if list is empty.
func (l *listInst[T]) First() Entry[T] {
	if l.IsNil() {
		return Entry[T]{}
	}

	return l.first
}

// Last returns the last entry, or nil if the list is empty.
func (l *listInst[T]) Last() Entry[T] {
	if l.IsNil() {
		return Entry[T]{}
	}

	return l.last
}

// Get returns entry with given index.
func (l *listInst[T]) Get(index int) Entry[T] {
	if l.IsNil() || index < 0 || index >= l.len {
		panic(fault.Trace(fault.ErrIndexOutOfRange))
	}

	entry := l.first
	for i := 0; i < index; i++ {
		entry = entry.next
	}

	return entry
}

// Prepend inserts new entry containing given value at front of list, and returns it.
func (l List[T]) Prepend(value T) Entry[T] {
	if l.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	entry := l.newEntry(value)
	l.putFirst(entry)

	l.len++

	return entry
}

// Append inserts new entry containing given value at back of list, and returns it.
func (l List[T]) Append(value T) Entry[T] {
	if l.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	entry := l.newEntry(value)
	l.putLast(entry)

	l.len++

	return entry
}

// InsertBefore inserts new entry containing given value before given pivot entry, and returns it.
func (l List[T]) InsertBefore(pivot Entry[T], value T) Entry[T] {
	if l.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if pivot.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if pivot.list != l {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	entry := l.newEntry(value)
	l.putBefore(pivot, entry)

	l.len++

	return entry
}

// InsertAfter inserts new entry containing given value after given pivot entry, and returns it.
func (l List[T]) InsertAfter(pivot Entry[T], value T) Entry[T] {
	if l.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if pivot.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if pivot.list != l {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	entry := l.newEntry(value)
	l.putAfter(pivot, entry)

	l.len++

	return entry
}

// Delete deletes given entry from list.
func (l List[T]) Delete(entry Entry[T]) {
	if l.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if entry.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if entry.list != l {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	l.cut(entry)
	entry.list = List[T]{}

	l.len--
}

// newEntry initializes and returns new entry of list.
func (l List[T]) newEntry(value T) Entry[T] {
	return Entry[T]{
		&entryInst[T]{
			list:  l,
			value: value,
		},
	}
}

// putFirst puts given new entry at front of list.
func (l List[T]) putFirst(entry Entry[T]) {
	if l.first.IsNil() {
		l.first = entry
		l.last = entry
		// entry.prev = nil
		// entry.next = nil
	} else {
		l.putBefore(l.first, entry)
	}
}

// putLast puts given new entry at back of list.
func (l List[T]) putLast(entry Entry[T]) {
	if l.last.IsNil() {
		l.putFirst(entry)
	} else {
		l.putAfter(l.last, entry)
	}
}

// putBefore puts given entry before given pivot entry.
func (l List[T]) putBefore(pivot, entry Entry[T]) {
	entry.next = pivot

	if pivot.prev.IsNil() {
		// entry.prev = nil
		l.first = entry
	} else {
		entry.prev = pivot.prev
		pivot.prev.next = entry
	}

	pivot.prev = entry
}

// putAfter puts given entry after given pivot entry.
func (l List[T]) putAfter(pivot, entry Entry[T]) {
	entry.prev = pivot

	if pivot.next.IsNil() {
		// entry.next = nil
		l.last = entry
	} else {
		entry.next = pivot.next
		pivot.next.prev = entry
	}

	pivot.next = entry
}

// cut cuts given entry from list.
func (l List[T]) cut(entry Entry[T]) {
	if entry.prev.IsNil() {
		l.first = entry.next
	} else {
		entry.prev.next = entry.next
	}

	if entry.next.IsNil() {
		l.last = entry.prev
	} else {
		entry.next.prev = entry.prev
	}

	entry.prev = Entry[T]{}
	entry.next = Entry[T]{}
}
