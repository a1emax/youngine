package orderedmap

import (
	"github.com/a1emax/youngine/extra/list"
)

// Entry is [OrderedMap] entry.
type Entry[K comparable, V any] struct {
	*entryInst[K, V]
}

// entryInst is the internal state of the [Entry] type (shared between copies).
//
// entryInst methods do not change it.
type entryInst[K comparable, V any] struct {
	listEntry list.Entry[Entry[K, V]]
	key       K
	value     V
}

// IsNil reports whether entry is nil.
func (e *entryInst[K, V]) IsNil() bool {
	return e == nil
}

// Prev returns the previous entry of the same map, or nil if there is no one.
func (e *entryInst[K, V]) Prev() Entry[K, V] {
	if e.IsNil() {
		return Entry[K, V]{}
	}

	listEntry := e.listEntry.Prev()
	if listEntry.IsNil() {
		return Entry[K, V]{}
	}

	return listEntry.Value()
}

// Next returns the next entry of the same map, or nil if there is no one.
func (e *entryInst[K, V]) Next() Entry[K, V] {
	if e.IsNil() {
		return Entry[K, V]{}
	}

	listEntry := e.listEntry.Next()
	if listEntry.IsNil() {
		return Entry[K, V]{}
	}

	return listEntry.Value()
}

// Key returns contained key.
func (e *entryInst[K, V]) Key() K {
	if e.IsNil() {
		var zero K
		return zero
	}

	return e.key
}

// Value returns contained value.
func (e *entryInst[K, V]) Value() V {
	if e.IsNil() {
		var zero V
		return zero
	}

	return e.value
}
