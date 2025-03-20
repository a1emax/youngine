package heap

import (
	"github.com/a1emax/youngine/fault"
)

// Core of heap on top of base of type B.
//
// Core algorithms are copied from [container/heap] package.
type Core[B Base[T], T any] struct {
	Base B
}

// NewCore initializes and returns new [Core].
func NewCore[B Base[T], T any](base B) Core[B, T] {
	return Core[B, T]{base}
}

// Len returns number of elements in heap.
func (c Core[B, T]) Len() int {
	// TODO: Return 0 if base is nil.

	return c.Base.Len()
}

// Restore establishes heap invariants required by other methods.
//
// Restore is idempotent with respect to invariants and may be called whenever they may have been invalidated.
func (c Core[B, T]) Restore() {
	n := c.Len()
	for i := n/2 - 1; i >= 0; i-- {
		c.down(i, n)
	}
}

// Insert inserts element with given value to heap.
func (c Core[B, T]) Insert(value T) {
	// TODO: Check base for nil.

	c.Base.Append(value)
	c.up(c.Base.Len() - 1)
}

// Update re-establishes heap ordering after element with given index has changed its value.
func (c Core[B, T]) Update(index int) {
	n := c.Len()
	if index < 0 || index >= n {
		panic(fault.Trace(fault.ErrIndexOutOfRange))
	}

	if !c.down(index, n) {
		c.up(index)
	}
}

// Extract deletes first element in order from heap and returns its value and true,
// or just returns zero value and false if heap is empty.
func (c Core[B, T]) Extract() (T, bool) {
	n := c.Len() - 1
	if n < 0 {
		var zero T

		return zero, false
	}

	c.Base.Swap(0, n)
	c.down(0, n)

	return c.Base.Pop(), true
}

// MustExtract calls Extract and panics if false is returned.
func (c Core[B, T]) MustExtract() T {
	value, ok := c.Extract()
	if !ok {
		panic(fault.Trace(fault.ErrInvalidUse))
	}

	return value
}

// Delete deletes element with given index from heap and returns its value.
func (c Core[B, T]) Delete(index int) T {
	n := c.Len() - 1
	if index < 0 || index > n {
		panic(fault.Trace(fault.ErrIndexOutOfRange))
	}

	if n != index {
		c.Base.Swap(index, n)
		if !c.down(index, n) {
			c.up(index)
		}
	}

	return c.Base.Pop()
}

// up re-establishes heap invariants by comparing and possibly swapping nodes with their parents.
func (c Core[B, T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !c.Base.Before(j, i) {
			break
		}
		c.Base.Swap(i, j)
		j = i
	}
}

// down re-establishes heap invariants by comparing and possibly swapping nodes with their children.
func (c Core[B, T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && c.Base.Before(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !c.Base.Before(j, i) {
			break
		}
		c.Base.Swap(i, j)
		i = j
	}

	return i > i0
}
