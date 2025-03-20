package textscroller

import (
	"math"

	"github.com/a1emax/youngine/basic"
)

// scrollObj represents scroll.
type scrollObj struct {
	isPrepared bool
	maxOffset  basic.Float
	offset     basic.Float
}

// prepareScroll resets scroll if element has just appeared on screen, or does nothing otherwise.
func (t *textScrollerImpl[T]) prepareScroll() {
	if t.scroll.isPrepared {
		return
	}

	t.scroll.isPrepared = true
	t.scroll.offset = 0.0
}

// excludeScroll handles element disappearing from screen.
func (t *textScrollerImpl[T]) excludeScroll() {
	t.scroll.isPrepared = false
}

// arrangeScroll recalculates scroll parameters according to text.
func (t *textScrollerImpl[T]) arrangeScroll() {
	if t.text.key.IsSet() {
		t.scroll.maxOffset = math.Floor(t.text.view.Size().Y() - t.bbox.Height())
	} else {
		t.scroll.maxOffset = 0.0
	}
}

// controlScroll returns scroll callbacks to be called from input controller.
func (t *textScrollerImpl[T]) controlScroll() (base, move func(position basic.Vec2)) {
	var prevY basic.Float
	base = func(position basic.Vec2) {
		prevY = position.Y()
	}
	move = func(position basic.Vec2) {
		velocity := prevY - position.Y()
		t.scroll.offset = math.Floor(t.scroll.offset + velocity)

		prevY = position.Y()
	}

	return base, move
}

// updateScroll updates scroll.
func (t *textScrollerImpl[T]) updateScroll() {
	t.scroll.offset = basic.Clamp(t.scroll.offset, 0.0, t.scroll.maxOffset)
}
