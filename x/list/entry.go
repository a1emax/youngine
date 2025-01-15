package list

// Entry of [List].
type Entry[T any] struct {
	*entryInst[T]
}

// entryInst is the internal state of the [Entry] type (shared between copies).
//
// entryInst methods do not change it.
type entryInst[T any] struct {
	list       List[T]
	prev, next Entry[T]
	value      T
}

// IsNil reports whether entry is nil.
func (e *entryInst[T]) IsNil() bool {
	return e == nil
}

// Prev returns the previous entry of the same list, or nil if there is no one.
func (e *entryInst[T]) Prev() Entry[T] {
	if e.IsNil() {
		return Entry[T]{}
	}

	return e.prev
}

// Next returns the next entry of the same list, or nil if there is no one.
func (e *entryInst[T]) Next() Entry[T] {
	if e.IsNil() {
		return Entry[T]{}
	}

	return e.next
}

// Value returns contained value.
func (e *entryInst[T]) Value() T {
	if e.IsNil() {
		var zero T
		return zero
	}

	return e.value
}
