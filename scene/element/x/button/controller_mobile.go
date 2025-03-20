//go:build android

package button

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input/element/touchscreentouch"
)

// initController initializes input controller.
func (b *buttonImpl[T]) initController(config Config) {
	b.controller = touchscreentouch.NewController(touchscreentouch.ControllerConfig[basic.None]{
		Clock: config.Clock,
		Input: config.Input.Touchscreen(),

		HitTest: func(position basic.Vec2) bool {
			return b.bbox.Contains(position)
		},

		OnHover: func(event touchscreentouch.HoverEvent[basic.None]) {
			b.isPressed = true

			if f := b.Props().OnPress; f != nil {
				f(PressEvent{
					Duration: event.Duration,
				})
			}
		},
		OnUp: func(event touchscreentouch.UpEvent) {
			b.isPressed = false

			if !event.JustEnded {
				return
			}

			if f := b.Props().OnClick; f != nil {
				f(ClickEvent{})
			}
		},
	})
}
