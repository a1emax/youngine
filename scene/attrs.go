package scene

import (
	"github.com/a1emax/youngine/basic"
)

// Attrs associated with element.
type Attrs struct {

	// MaxHeight specifies maximum height of element.
	MaxHeight basic.Opt[basic.Float]

	// MaxWidth specifies maximum width of element.
	MaxWidth basic.Opt[basic.Float]

	// MinHeight specifies minimum height of element.
	MinHeight basic.Opt[basic.Float]

	// MinWidth specifies minimum width of element.
	MinWidth basic.Opt[basic.Float]

	// PreHeight specifies preliminary height of element.
	PreHeight basic.Opt[basic.Float]

	// PreWidth specifies preliminary width of element.
	PreWidth basic.Opt[basic.Float]
}

// FixHeight returns outline with fixed element height.
func (a Attrs) FixHeight(value basic.Float) Attrs {
	a.MaxHeight = basic.SetOpt(value)
	a.MinHeight = basic.SetOpt(value)
	a.PreHeight = basic.SetOpt(value)

	return a
}

// FixWidth returns outline with fixed element width.
func (a Attrs) FixWidth(value basic.Float) Attrs {
	a.MaxWidth = basic.SetOpt(value)
	a.MinWidth = basic.SetOpt(value)
	a.PreWidth = basic.SetOpt(value)

	return a
}
