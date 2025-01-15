package input

// Keyboard state.
type Keyboard interface {

	// Key returns [KeyboardKey] with given code, or panics if code is unknown.
	Key(code KeyboardKeyCode) KeyboardKey

	// Mark marks state until the next update.
	Mark()
}
