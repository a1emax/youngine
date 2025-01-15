package input

// Touchscreen state.
type Touchscreen interface {

	// TouchCount returns number of touches.
	//
	// NOTE that result remains valid only until the next update.
	TouchCount() int

	// Touch returns [TouchscreenTouch] with given index, or panics if index is out of range.
	//
	// NOTE that result remains valid only until the next update.
	Touch(index int) TouchscreenTouch

	// Mark marks state until the next update.
	Mark()
}
