package input

// MouseButtonCode identifies mouse button.
type MouseButtonCode int

const (
	_ MouseButtonCode = iota
	MouseButtonCodeLeft
	MouseButtonCodeMiddle
	MouseButtonCodeRight
	_
	_

	// MaxMouseButtonCode specifies the maximum known value of the [MouseButtonCode] type.
	MaxMouseButtonCode MouseButtonCode = iota - 1

	// MinMouseButtonCode specifies the minimum known value of the [MouseButtonCode] type.
	MinMouseButtonCode MouseButtonCode = 1

	// MouseButtonCount specifies the number of known mouse buttons.
	MouseButtonCount = int(MaxMouseButtonCode)
)
