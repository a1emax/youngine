package clock

// Clock of logical time.
type Clock interface {

	// Now returns current logical time. It is guaranteed that result is greater than previous one
	// after the next update start and is equal to previous one during one update.
	Now() Time
}
