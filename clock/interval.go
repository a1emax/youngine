package clock

import (
	"github.com/a1emax/youngine/fault"
)

// CheckInterval reports whether number of ticks, including current one, that have passed since given logical time,
// is multiple of given positive interval.
//
// CheckInterval always returns false if time is zero.
func CheckInterval(clk Clock, since Time, interval Ticks) bool {
	if clk == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if interval <= 0 {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	if since.IsZero() {
		return false
	}

	d := clk.Now().Sub(since) + 1

	return d >= interval && d%interval == 0
}
