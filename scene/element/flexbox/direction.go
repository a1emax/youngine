package flexbox

// Direction is the type of the [Config.Direction] property.
type Direction int

const (
	// DirectionRow specifies that container's main axis is horizontal and directed from left to right.
	DirectionRow Direction = iota

	// DirectionColumn specifies that container's main axis is vertical and directed from top to bottom.
	DirectionColumn

	// maxDirection specifies the maximum valid value of the [Direction] type.
	maxDirection Direction = iota - 1

	// minDirection specifies the minimum valid value of the [Direction] type.
	minDirection Direction = 0
)

// IsValid reports whether value is valid.
func (d Direction) IsValid() bool {
	return d >= minDirection && d <= maxDirection
}

// IsDefault reports whether value is default or invalid.
func (d Direction) IsDefault() bool {
	return d == minDirection || !d.IsValid()
}
