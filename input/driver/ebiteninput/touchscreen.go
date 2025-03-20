package ebiteninput

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// Touchscreen state based on Ebitengine.
type Touchscreen interface {
	input.Touchscreen

	// Update updates state.
	Update()
}

// touchscreenImpl is the implementation of the [Touchscreen] interface.
type touchscreenImpl struct {
	helper  touchscreenHelper
	touches []touchscreenTouchImpl
}

// NewTouchscreen initializes and returns new [Touchscreen].
func NewTouchscreen(clk clock.Clock) Touchscreen {
	if clk == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	t := &touchscreenImpl{}
	t.init(clk)

	return t
}

// init initializes [Touchscreen].
func (t *touchscreenImpl) init(clk clock.Clock) {
	t.helper = newTouchscreenHelper(clk)
}

// TouchCount implements the [input.Touchscreen] interface.
func (t *touchscreenImpl) TouchCount() int {
	return len(t.touches)
}

// Touch implements the [input.Touchscreen] interface.
func (t *touchscreenImpl) Touch(index int) input.TouchscreenTouch {
	if index < 0 || index > len(t.touches) {
		panic(fault.Trace(fault.ErrIndexOutOfRange))
	}

	return &t.touches[index]
}

// Mark implements the [input.Touchscreen] interface.
func (t *touchscreenImpl) Mark() {
	for i := range t.touches {
		t.touches[i].Mark()
	}
}

// Update implements the [Touchscreen] interface.
func (t *touchscreenImpl) Update() {
	t.helper.update()

	t.touches = t.touches[:0]
	for _, touch := range t.helper.touches {
		if touch.exposedAt.IsZero() {
			continue
		}

		t.touches = append(t.touches, touchscreenTouchImpl{
			id:        input.TouchscreenTouchID(touch.id),
			startedAt: touch.exposedAt,
			position:  basic.Vec2{basic.Float(touch.x), basic.Float(touch.y)},
			isMarked:  false,
		})
	}
}
