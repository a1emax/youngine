package ebiteninput

import (
	"github.com/a1emax/youngine/clock"
)

type testClockFunc func() clock.Time

func (f testClockFunc) Now() clock.Time {
	return f()
}
