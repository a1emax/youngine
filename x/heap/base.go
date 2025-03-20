package heap

// Base represents array of elements with values of type T.
//
// Base implementations are intended to be used only by [Core] and code that uses it locally,
// so they are usually simplified, e.g. do not validate arguments.
type Base[T any] interface {

	// Len returns number of elements.
	Len() int

	// Before reports whether ith element comes before jth one in logical order.
	Before(i, j int) bool

	// Swap swaps ith and jth elements.
	Swap(i, j int)

	// Append appends element with value x to array's end.
	Append(x T)

	// Pop deletes last element from array and returns it's value.
	Pop() T
}
