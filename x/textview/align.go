package textview

// Align is text alignment during drawing.
type Align int

const (

	// AlignLeft specifies that text is aligned along left margin.
	AlignLeft Align = iota

	// AlignRight specifies that text is aligned along right margin.
	AlignRight

	// AlignCenter specifies that text is aligned to neither left nor right margin; there is even gap
	// on each side of each line.
	AlignCenter

	// AlignJustify specifies that text is aligned along left margin, with word-spacing adjusted so that
	// text falls flush with both margins.
	AlignJustify

	// maxAlign specifies the maximum valid value of the [Align] type.
	maxAlign Align = iota - 1

	// minAlign specifies the minimum valid value of the [Align] type.
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
