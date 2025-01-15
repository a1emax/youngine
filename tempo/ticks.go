package tempo

// Ticks represents number of logical time units. One tick usually corresponds to single application update.
type Ticks int64

// Tick represents single tick.
const Tick Ticks = 1

// Second represents the number of ticks per second. This is the default value - use your own for custom cases.
const Second = 60 * Tick

// Minute represents the number of ticks per minute. This is the default value - use your own for custom cases.
const Minute = 60 * Second

// Hour represents the number of ticks per hour. This is the default value - use your own for custom cases.
const Hour = 60 * Minute
