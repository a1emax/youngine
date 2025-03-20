package steadymap

import (
	"iter"

	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/x/list"
)

// SteadyMap is map with keys of type K arranged in insertion order and associated with values of type V.
//
// Values of this type are references to shared internal state. Use the Copy method to make separate copy.
type SteadyMap[K comparable, V any] struct {
	hash map[K]*steadyMapEntry[K, V]
	list list.List[*steadyMapEntry[K, V]]
}

// steadyMapEntry is internal representation of map key.
type steadyMapEntry[K comparable, V any] struct {
	marker list.Marker[*steadyMapEntry[K, V]]
	key    K
	value  V
}

// New initializes and returns new [SteadyMap].
func New[K comparable, V any]() SteadyMap[K, V] {
	return SteadyMap[K, V]{
		hash: make(map[K]*steadyMapEntry[K, V]),
		list: list.New[*steadyMapEntry[K, V]](),
	}
}

// IsNil reports whether map is nil.
func (m SteadyMap[K, V]) IsNil() bool {
	return m.hash == nil
}

// Copy returns copy of map.
func (m SteadyMap[K, V]) Copy() SteadyMap[K, V] {
	if m.IsNil() {
		return SteadyMap[K, V]{}
	}

	result := New[K, V]()

	for entry := range m.list.All() {
		resultEntry := &steadyMapEntry[K, V]{
			key:   entry.key,
			value: entry.value,
		}

		result.hash[resultEntry.key] = resultEntry
		resultEntry.marker = result.list.Append(resultEntry)
	}

	return result
}

// All returns iterator that iterates over keys in direct order, producing them and associated values.
func (m SteadyMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if m.IsNil() {
			return
		}

		for entry := range m.list.All() {
			if !yield(entry.key, entry.value) {
				break
			}
		}
	}
}

// Backward returns iterator that iterates over keys in backward order, producing them and associated values.
func (m SteadyMap[K, V]) Backward() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		if m.IsNil() {
			return
		}

		for entry := range m.list.Backward() {
			if !yield(entry.key, entry.value) {
				break
			}
		}
	}
}

// Len returns number of keys.
func (m SteadyMap[K, V]) Len() int {
	if m.IsNil() {
		return 0
	}

	return len(m.hash)
}

// Has reports whether map contains given key.
func (m SteadyMap[K, V]) Has(key K) bool {
	if m.IsNil() {
		return false
	}

	_, ok := m.hash[key]

	return ok
}

// Get returns value associated with given key, or zero value if map does not contain key.
func (m SteadyMap[K, V]) Get(key K) V {
	if !m.IsNil() {
		if entry, ok := m.hash[key]; ok {
			return entry.value
		}
	}

	var zero V

	return zero
}

// Set associates given value with given key. If map does not contain key, it will be inserted at back of map.
func (m SteadyMap[K, V]) Set(key K, value V) {
	if m.IsNil() {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	entry, ok := m.hash[key]
	if !ok {
		entry = &steadyMapEntry[K, V]{
			key: key,
		}

		m.hash[key] = entry
		entry.marker = m.list.Append(entry)
	}

	entry.value = value
}

// Delete deletes given key from map, if it contains one.
func (m SteadyMap[K, V]) Delete(key K) {
	if m.IsNil() {
		return // Nothing needs to be done.
	}

	entry, ok := m.hash[key]
	if !ok {
		return // Nothing needs to be done.
	}

	delete(m.hash, key)
	m.list.Delete(entry.marker)
	// *entry = steadyMapEntry[K, V]{}
}
