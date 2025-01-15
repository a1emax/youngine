package tempo

// Clock controls logical time.
type Clock interface {
	Nower

	// Update updates clock.
	Update()
}

// clockImpl is the implementation of the [Clock] interface.
type clockImpl struct {
	ticks Ticks
}

// NewClock initializes and returns new [Clock].
func NewClock() Clock {
	return &clockImpl{}
}

// Now implements the [Nower] interface.
func (c *clockImpl) Now() Time {
	return Time{c.ticks}
}

// Update implements the [Clock] interface.
func (c *clockImpl) Update() {
	c.ticks++
}
