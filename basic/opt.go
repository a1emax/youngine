package basic

import (
	"fmt"
)

// Opt is optional value of type T.
type Opt[T any] struct {
	isSet bool
	value T
}

// SetOpt returns set optional value.
func SetOpt[T any](value T) Opt[T] {
	return Opt[T]{true, value}
}

// String implements the [fmt.Stringer] interface.
func (o Opt[T]) String() string {
	if !o.isSet {
		return "<unset>"
	}

	return fmt.Sprintf("%v", o.value)
}

// IsSet reports whether value is set.
func (o Opt[T]) IsSet() bool {
	return o.isSet
}

// Get returns value if it is set, or zero value otherwise.
func (o Opt[T]) Get() T {
	return o.value
}

// Or returns value if it is set, or given value otherwise.
func (o Opt[T]) Or(value T) T {
	if o.isSet {
		return o.value
	}

	return value
}
