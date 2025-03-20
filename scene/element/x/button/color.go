package button

import (
	"image/color"
)

// color returns first given color is button is pressed, or second one otherwise.
func (b *buttonImpl[T]) color(primary, pressed color.Color) color.Color {
	if b.isPressed {
		return pressed
	}

	return primary
}
