//go:build windows

package textscroller

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/input/element/mousebutton"
	"github.com/a1emax/youngine/input/element/mousecursor"
)

// initController initializes input controller.
func (t *textScrollerImpl[T]) initController(config Config) {
	baseScroll, moveScroll := t.controlScroll()

	t.controller = mousecursor.NewController(mousecursor.ControllerConfig[basic.None]{
		Clock: config.Clock,
		Input: config.Input.Mouse().Cursor(),

		HitTest: func(position basic.Vec2) bool {
			return t.bbox.Contains(position)
		},

		Slave: mousebutton.NewController(mousebutton.ControllerConfig[mousecursor.Background[basic.None]]{
			Clock: config.Clock,
			Input: config.Input.Mouse().Button(input.MouseButtonCodeLeft),

			OnDown: func(event mousebutton.DownEvent[mousecursor.Background[basic.None]]) {
				baseScroll(event.Background.Position)
			},
			OnPress: func(event mousebutton.PressEvent[mousecursor.Background[basic.None]]) {
				moveScroll(event.Background.Position)
			},
		}),
	})
}
