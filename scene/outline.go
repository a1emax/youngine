package scene

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
)

// Outline of element.
type Outline struct {

	// MinHeight specifies minimum height of element.
	MinHeight basic.Opt[basic.Float]

	// MinWidth specifies minimum width of element.
	MinWidth basic.Opt[basic.Float]

	// MaxHeight specifies maximum height of element.
	MaxHeight basic.Opt[basic.Float]

	// MaxWidth specifies maximum width of element.
	MaxWidth basic.Opt[basic.Float]

	// PreHeight specifies preliminary height of element.
	PreHeight basic.Opt[basic.Float]

	// PreWidth specifies preliminary width of element.
	PreWidth basic.Opt[basic.Float]
}

// SetWidth sets width of element to given value.
func (o *Outline) SetWidth(value basic.Float) {
	if o == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	o.MinWidth = basic.SetOpt(value)
	o.MaxWidth = basic.SetOpt(value)
	o.PreWidth = basic.SetOpt(value)
}

// SetHeight sets height of element to given value.
func (o *Outline) SetHeight(value basic.Float) {
	if o == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	o.MinHeight = basic.SetOpt(value)
	o.MaxHeight = basic.SetOpt(value)
	o.PreHeight = basic.SetOpt(value)
}
