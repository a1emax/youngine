package ebiteninput

import (
	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
)

// Mouse state based on Ebitengine.
type Mouse interface {
	input.Mouse

	// Update updates state.
	Update()
}

// mouseImpl is the implementation of the [Mouse] interface.
type mouseImpl struct {
	buttons [input.MouseButtonCount]mouseButtonImpl
	cursor  mouseCursorImpl
	wheel   mouseWheelImpl
}

// NewMouse initializes and returns new [Mouse].
func NewMouse(clk clock.Clock) Mouse {
	if clk == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	m := &mouseImpl{}
	m.init(clk)

	return m
}

// init initializes [Mouse].
func (m *mouseImpl) init(clk clock.Clock) {
	for i := range m.buttons {
		m.buttons[i].init(input.MouseButtonCode(i+1), clk)
	}
	m.cursor.init()
	m.wheel.init()
}

// Button implements the [input.Mouse] interface.
func (m *mouseImpl) Button(code input.MouseButtonCode) input.MouseButton {
	if code < input.MinMouseButtonCode || code > input.MaxMouseButtonCode {
		panic(fault.Trace(fault.ErrInvalidArgument))
	}

	return &m.buttons[code-1]
}

// Cursor implements the [input.Mouse] interface.
func (m *mouseImpl) Cursor() input.MouseCursor {
	return &m.cursor
}

// Wheel implements the [input.Mouse] interface.
func (m *mouseImpl) Wheel() input.MouseWheel {
	return &m.wheel
}

// Mark implements the [input.Mouse] interface.
func (m *mouseImpl) Mark() {
	for i := range m.buttons {
		m.buttons[i].Mark()
	}
	m.cursor.Mark()
	m.wheel.Mark()
}

// Update implements the [Mouse] interface.
func (m *mouseImpl) Update() {
	for i := range m.buttons {
		m.buttons[i].Update()
	}
	m.cursor.Update()
	m.wheel.Update()
}
