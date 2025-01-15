package input

// Mouse state.
type Mouse interface {

	// Button returns [MouseButton] with given code, or panics if code is unknown.
	Button(code MouseButtonCode) MouseButton

	// Cursor returns [MouseCursor].
	Cursor() MouseCursor

	// Wheel returns [MouseWheel].
	Wheel() MouseWheel

	// Mark marks state until the next update.
	Mark()
}
