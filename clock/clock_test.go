package clock

type testClockFunc func() Time

func (c testClockFunc) Now() Time {
	return c()
}
