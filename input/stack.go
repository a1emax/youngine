package input

// Stack state.
type Stack interface {

	// Keyboard returns [Keyboard].
	Keyboard() Keyboard

	// Mouse returns [Mouse].
	Mouse() Mouse

	// Touchscreen returns [Touchscreen].
	Touchscreen() Touchscreen

	// Mark marks state until the next update.
	Mark()
}
