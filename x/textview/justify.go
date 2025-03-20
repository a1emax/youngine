package textview

// Justify is justification of items along container's axis (horizontal for words, vertical for lines).
type Justify int

const (

	// JustifyStart specifies that items are packed flush to each other toward container's start edge.
	JustifyStart Justify = iota

	// JustifyCenter specifies that items are packed flush to each other toward container's center.
	JustifyCenter

	// JustifyEnd specifies that items are packed flush to each other toward container's end edge.
	JustifyEnd

	// JustifySpaceBetween specifies that items are evenly distributed within container. Spacing between each pair
	// of adjacent items is the same. The first item is flush with container's start edge, and the last item is
	// flush with container's end edge.
	JustifySpaceBetween

	// JustifySpaceAround specifies that items are evenly distributed within container. Spacing between each pair
	// of adjacent items is the same. Empty space before the first and after the last item equals half of space
	// between each pair of adjacent items.
	JustifySpaceAround

	// JustifySpaceEvenly specifies that items are evenly distributed within container. Spacing between each pair
	// of adjacent items, container's start edge and the first item, and container's end edge and the last item,
	// are all exactly the same.
	JustifySpaceEvenly

	// maxJustify specifies the maximum valid value of the [Justify] type.
	maxJustify Justify = iota - 1

	// minJustify specifies the minimum valid value of the [Justify] type.
	minJustify Justify = 0
)

// IsValid reports whether value is valid.
func (j Justify) IsValid() bool {
	return j >= minJustify && j <= maxJustify
}
