package flexbox

// Align is alignment of item on container's cross axis.
type Align int

const (

	// AlignStretch specifies that item is stretched such that its size is same as container while respecting
	// width and height constraints.
	AlignStretch Align = iota

	// AlignStart specifies that item's start edge is flushed with container's start edge.
	AlignStart

	// AlignCenter specifies that item is centered within container. If item is larger than container, it will
	// overflow equally in both directions.
	AlignCenter

	// AlignEnd specifies that item's end edge is flushed with container's end edge.
	AlignEnd

	// maxAlign specifies the maximum valid value of the [Align] type.
	maxAlign Align = iota - 1

	// minAlign specifies the minimum valid value if the [Align] type.
	minAlign Align = 0
)

// IsValid reports whether value is valid.
func (a Align) IsValid() bool {
	return a >= minAlign && a <= maxAlign
}

// IsDefault reports whether value is default or invalid.
func (a Align) IsDefault() bool {
	return a == minAlign || !a.IsValid()
}
