package orderedmap

import (
	"github.com/a1emax/youngine/extra/list"
	"github.com/a1emax/youngine/fault"
)

// OrderedMap is mapping of keys of type K to values of type V that preserves order of entries.
type OrderedMap[K comparable, V any] struct {
	*orderedMapInst[K, V]
}

// ReadOnly is [OrderedMap] with read-only access.
type ReadOnly[K comparable, V any] struct {
	*orderedMapInst[K, V]
}

// orderedMapInst is the internal state of the [OrderedMap] type (shared between copies).
//
// orderedMapInst methods do not change it.
type orderedMapInst[K comparable, V any] struct {
	hashMap map[K]Entry[K, V]
	list    list.List[Entry[K, V]]
}

// New initializes and returns new [OrderedMap].
func New[K comparable, V any]() OrderedMap[K, V] {
	m := OrderedMap[K, V]{
		&orderedMapInst[K, V]{},
	}

	m.hashMap = make(map[K]Entry[K, V])
	m.list = list.New[Entry[K, V]]()

	return m
}

// ReadOnly returns map with read-only access.
func (m OrderedMap[K, V]) ReadOnly() ReadOnly[K, V] {
	return ReadOnly[K, V](m)
}

// Copy returns copy of map.
func (m *orderedMapInst[K, V]) Copy() OrderedMap[K, V] {
	if m.IsNil() {
		return OrderedMap[K, V]{}
	}

	result := New[K, V]()

	for e := m.list.First(); !e.IsNil(); e = e.Next() {
		entry := e.Value()
		resultEntry := Entry[K, V]{
			&entryInst[K, V]{
				key:   entry.key,
				value: entry.value,
			},
		}

		result.hashMap[resultEntry.key] = resultEntry
		resultEntry.listEntry = result.list.Append(resultEntry)
	}

	return result
}

// IsNil reports whether  map is nil.
func (m *orderedMapInst[K, V]) IsNil() bool {
	return m == nil
}

// Len returns number of entries.
func (m *orderedMapInst[K, V]) Len() int {
	if m.IsNil() {
		return 0
	}

	return len(m.hashMap)
}

// First returns the first entry, or nil if map is empty.
func (m *orderedMapInst[K, V]) First() Entry[K, V] {
	if m.IsNil() {
		return Entry[K, V]{}
	}

	return m.list.First().Value()
}

// Last returns the last entry, or nil if map is empty.
func (m *orderedMapInst[K, V]) Last() Entry[K, V] {
	if m.IsNil() {
		return Entry[K, V]{}
	}

	return m.list.Last().Value()
}

// Contains reports whether map contains given key.
func (m *orderedMapInst[K, V]) Contains(key K) bool {
	if m.IsNil() {
		return false
	}

	_, ok := m.hashMap[key]

	return ok
}

// Get returns value associated with given key, or zero value if map does not contain key.
func (m *orderedMapInst[K, V]) Get(key K) V {
	if !m.IsNil() {
		if entry, ok := m.hashMap[key]; ok {
			return entry.value
		}
	}

	var zero V

	return zero
}

// Set associates given value with given key. If map does not contain key, new entry containing one and value
// will be inserted at back of map.
func (m OrderedMap[K, V]) Set(key K, value V) {
	if m.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	entry, ok := m.hashMap[key]
	if !ok {
		entry = Entry[K, V]{
			&entryInst[K, V]{
				key: key,
			},
		}

		m.hashMap[key] = entry
		entry.listEntry = m.list.Append(entry)
	}

	entry.value = value
}

// Delete deletes given key from map, if it contains one.
func (m OrderedMap[K, V]) Delete(key K) {
	if m.IsNil() {
		return // Nothing needs to be done.
	}

	entry, ok := m.hashMap[key]
	if !ok {
		return // Nothing needs to be done.
	}

	delete(m.hashMap, key)
	m.list.Delete(entry.listEntry)
	entry.listEntry = list.Entry[Entry[K, V]]{}
}
