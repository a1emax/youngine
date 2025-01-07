package audio

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// Volumer provides volume.
type Volumer interface {

	// Volume returns volume.
	Volume() basic.Float
}

// VolumerFunc is the functional implementation of the [Volumer] interface.
type VolumerFunc func() basic.Float

// Volume implements the [Volumer] interface.
func (f VolumerFunc) Volume() basic.Float {
	if f == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return f()
}
