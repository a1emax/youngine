package scene

// Element of scene placed on screen of type S inside region of type R.
//
// Element can be in one of following statuses:
//   - Operated: element is displayed and interacts;
//   - Visible: element is displayed, but does not interact;
//   - Hidden: element is not displayed and does not interact.
type Element[S any, R Region] interface {

	// Region returns region element is placed inside.
	//
	// NOTE that bounding rectangle of region is valid only starting from arrange stage.
	Region() R

	// IsActive reports whether element is active.
	//
	// NOTE that result is valid only after refresh stage.
	//
	// NOTE that inactive element can only be hidden.
	IsActive() bool

	// Outline returns outline of element.
	//
	// NOTE that result is valid only after refresh stage.
	Outline() Outline

	// Refresh refreshes element, e.g. by recalculating its state.
	//
	// NOTE that containers should refresh all items at this stage, regardless of their statuses.
	Refresh()

	// Prepare prepares element, e.g. by precalculating some parameters.
	//
	// NOTE that containers should prepare operated and visible items and exclude hidden ones at this stage.
	Prepare()

	// Exclude excludes element, e.g. by cleaning up some resources for idle time.
	//
	// NOTE that containers should exclude all items at this stage, regardless of their statuses.
	Exclude()

	// Arrange arranges element, e.g. by executing layout algorithm.
	//
	// NOTE that containers should arrange operated and visible items at this stage.
	Arrange()

	// Actuate actuates element, e.g. by handling input.
	//
	// NOTE that containers should actuate operated items and inhibit visible and hidden ones at this stage.
	//
	// NOTE that containers should iterate through items in reverse order at this stage.
	Actuate()

	// Inhibit inhibits element, e.g. by ignoring input.
	//
	// NOTE that containers should inhibit all items at this stage, regardless of their statuses.
	//
	// NOTE that containers should iterate through items in reverse order at this stage.
	Inhibit()

	// Update updates element.
	//
	// NOTE that containers should update operated and visible items at this stage.
	Update()

	// Draw draws element on given screen.
	//
	// NOTE that containers should draw operated and visible items at this stage.
	Draw(screen S)

	// Dispose disposes element.
	//
	// NOTE that containers should dispose all items at this stage, regardless of their statuses.
	Dispose()
}

// BaseElement is the base implementation of the [Element] interface.
type BaseElement[S any, R Region] struct{}

// Interface compatibility check.
var _ Element[any, Region] = BaseElement[any, Region]{}

// Region implements the [Element] interface.
//
// NOTE that this method should be overloaded in most cases to return region passed to element initializer.
func (BaseElement[S, R]) Region() R {
	var zero R

	return zero
}

// IsActive implements the [Element] interface.
func (BaseElement[S, R]) IsActive() bool {
	return true
}

// Outline implements the [Element] interface.
func (BaseElement[S, R]) Outline() Outline {
	return Outline{}
}

// Refresh implements the [Element] interface.
func (BaseElement[S, R]) Refresh() {
}

// Prepare implements the [Element] interface.
func (BaseElement[S, R]) Prepare() {
}

// Exclude implements the [Element] interface.
func (BaseElement[S, R]) Exclude() {
}

// Arrange implements the [Element] interface.
func (BaseElement[S, R]) Arrange() {
}

// Actuate implements the [Element] interface.
func (BaseElement[S, R]) Actuate() {
}

// Inhibit implements the [Element] interface.
func (BaseElement[S, R]) Inhibit() {
}

// Update implements the [Element] interface.
func (BaseElement[S, R]) Update() {
}

// Draw implements the [Element] interface.
func (BaseElement[S, R]) Draw(screen S) {
}

// Dispose implements the [Element] interface.
func (BaseElement[S, R]) Dispose() {
}
