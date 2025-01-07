package tempo

import (
	"github.com/a1emax/youngine/fault"
)

// CheckInterval reports whether number of ticks, including current one, that have passed since given logical time,
// is multiple of given positive interval.
//
// CheckInterval always returns false if time is zero.
func CheckInterval(nower Nower, since Time, interval Ticks) bool {
	if nower == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if interval <= 0 {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	if since.IsZero() {
		return false
	}

	d := nower.Now().Sub(since) + 1

	return d >= interval && d%interval == 0
}
