package tempo

import (
	"fmt"
)

// Time represents number of ticks, including the first one, that have passed since start of logical time.
//
// Zero Time precedes start of logical time and represents both zero and any negative number of ticks.
type Time struct {
	ticks Ticks
}

// At returns [Time] representing given number of ticks.
func At(ticks Ticks) Time {
	return Time{max(0, ticks)}
}

// String implements the [fmt.Stringer] interface.
func (t Time) String() string {
	return fmt.Sprintf("%d", t.ticks)
}

// Ticks returns number of ticks represented by t.
func (t Time) Ticks() Ticks {
	return t.ticks
}

// IsZero reports whether t is zero.
func (t Time) IsZero() bool {
	return t.ticks == 0
}

// Add returns sum of t and d.
func (t Time) Add(d Ticks) Time {
	return Time{max(0, t.ticks+d)}
}

// Sub returns difference of t and u.
func (t Time) Sub(u Time) Ticks {
	return t.ticks - u.ticks
}
