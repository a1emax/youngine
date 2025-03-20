//go:build windows

package button

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/input/element/mousebutton"
	"github.com/a1emax/youngine/input/element/mousecursor"
)

// initController initializes input controller.
func (b *buttonImpl[T]) initController(config Config) {
	b.controller = mousecursor.NewController(mousecursor.ControllerConfig[basic.None]{
		Clock: config.Clock,
		Input: config.Input.Mouse().Cursor(),

		HitTest: func(position basic.Vec2) bool {
			return b.bbox.Contains(position)
		},

		Slave: mousebutton.NewController(mousebutton.ControllerConfig[mousecursor.Background[basic.None]]{
			Clock: config.Clock,
			Input: config.Input.Mouse().Button(input.MouseButtonCodeLeft),

			OnPress: func(event mousebutton.PressEvent[mousecursor.Background[basic.None]]) {
				b.isPressed = true

				if f := b.Props().OnPress; f != nil {
					f(PressEvent{
						Duration: event.Duration,
					})
				}
			},
			OnUp: func(event mousebutton.UpEvent) {
				b.isPressed = false

				if !event.JustReleased {
					return
				}

				if f := b.Props().OnClick; f != nil {
					f(ClickEvent{})
				}
			},
		}),
	})
}
