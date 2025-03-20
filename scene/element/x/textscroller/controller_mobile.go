//go:build android

package textscroller

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input/element/touchscreentouch"
)

// initController initializes input controller.
func (t *textScrollerImpl[T]) initController(config Config) {
	baseScroll, moveScroll := t.controlScroll()

	t.controller = touchscreentouch.NewController(touchscreentouch.ControllerConfig[basic.None]{
		Clock: config.Clock,
		Input: config.Input.Touchscreen(),

		HitTest: func(position basic.Vec2) bool {
			return t.bbox.Contains(position)
		},

		OnDown: func(event touchscreentouch.DownEvent[basic.None]) {
			baseScroll(event.Position)
		},
		OnHover: func(event touchscreentouch.HoverEvent[basic.None]) {
			moveScroll(event.Position)
		},
	})
}
