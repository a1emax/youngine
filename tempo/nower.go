package tempo

import (
	"github.com/a1emax/youngine/fault"
)

// Nower provides current logical time.
type Nower interface {

	// Now returns current logical time.
	Now() Time
}

// NowerFunc is the functional implementation of the [Nower] interface.
type NowerFunc func() Time

// Now implements the [Nower] interface.
func (f NowerFunc) Now() Time {
	if f == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return f()
}
