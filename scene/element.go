package scene

import (
	"github.com/a1emax/youngine/basic"
)

// Element of scene displayed on screen of type S and extended by trait of type T.
//
// Element can be in one of following statuses:
//   - Interactive: element is displayed and interacts;
//   - Visible: element is displayed, but may not interact (interactive elements are visible);
//   - Hidden: element is not displayed and does not interact (off elements are hidden).
type Element[S, T any] interface {

	// Trait returns trait of element.
	//
	// Trait result is valid only after refresh stage.
	Trait() T

	// IsOff reports whether element is off.
	//
	// IsOff result is valid only after refresh stage.
	IsOff() bool

	// Attrs returns attrs of element.
	//
	// Attrs result is valid only after prepare stage.
	Attrs() Attrs

	// Refresh refreshes element, e.g. by recalculating its trait and props.
	//
	// Containers should refresh all items at this stage, regardless of their statuses.
	Refresh()

	// Prepare prepares element for this iteration, e.g. by precalculating some parameters.
	//
	// Containers should prepare visible items and exclude hidden ones at this stage.
	Prepare()

	// Exclude excludes element for this iteration, e.g. by cleaning up some resources for idle time.
	//
	// Containers should exclude all items at this stage, regardless of their statuses.
	Exclude()

	// Arrange arranges element content within given bounding box, e.g. by executing layout algorithm.
	//
	// Containers should arrange visible items at this stage.
	Arrange(bbox basic.Rect)

	// Actuate actuates element, e.g. by handling input.
	//
	// Containers should actuate interactive items and inhibit other ones at this stage.
	//
	// Containers should iterate through items in reverse order at this stage.
	Actuate()

	// Inhibit inhibits element, e.g. by ignoring input.
	//
	// Containers should inhibit all items at this stage, regardless of their statuses.
	//
	// Containers should iterate through items in reverse order at this stage.
	Inhibit()

	// Update updates element.
	//
	// Containers should update visible items at this stage.
	Update()

	// Draw draws element on given screen.
	//
	// Containers should draw visible items at this stage.
	Draw(screen S)

	// Dispose disposes element.
	//
	// Containers should dispose all items at this stage, regardless of their statuses.
	Dispose()
}

// BaseElement is the base implementation of the [Element] interface with props of type P.
//
// BaseElement is intended for internal use only as part of real element.
type BaseElement[S, T, P any] struct {
	trait     T
	traitFunc TraitFunc[T]
	props     P
	propsFunc PropsFunc[P]
}

// TraitFunc accepts current trait of type T and returns new one.
type TraitFunc[T any] func(T) T

// PropsFunc accepts current props of type P and returns new ones.
type PropsFunc[P any] func(P) P

// Interface compatibility check.
var _ Element[any, any] = &BaseElement[any, any, any]{}

// Init initializes element.
//
// Init does not validate input parameters.
func (e *BaseElement[S, T, P]) Init(traitFunc TraitFunc[T], propsFunc PropsFunc[P]) {
	e.traitFunc = traitFunc
	e.propsFunc = propsFunc
}

// Trait implements the [Element] interface.
func (e *BaseElement[S, T, P]) Trait() T {
	return e.trait
}

// Props returns element props.
//
// Props result is valid only after refresh stage.
func (e *BaseElement[S, T, P]) Props() P {
	return e.props
}

// IsOff implements the [Element] interface.
//
// This method should only be overloaded if element can be off internally.
func (e *BaseElement[S, T, P]) IsOff() bool {
	return false
}

// Attrs implements the [Element] interface.
func (e *BaseElement[S, T, P]) Attrs() Attrs {
	return Attrs{}
}

// Refresh implements the [Element] interface.
func (e *BaseElement[S, T, P]) Refresh() {
	e.trait = e.traitFunc(e.trait)
	e.props = e.propsFunc(e.props)
}

// Prepare implements the [Element] interface.
func (e *BaseElement[S, T, P]) Prepare() {
}

// Exclude implements the [Element] interface.
func (e *BaseElement[S, T, P]) Exclude() {
}

// Arrange implements the [Element] interface.
func (e *BaseElement[S, T, P]) Arrange(bbox basic.Rect) {
}

// Actuate implements the [Element] interface.
func (e *BaseElement[S, T, P]) Actuate() {
}

// Inhibit implements the [Element] interface.
func (e *BaseElement[S, T, P]) Inhibit() {
}

// Update implements the [Element] interface.
func (e *BaseElement[S, T, P]) Update() {
}

// Draw implements the [Element] interface.
func (e *BaseElement[S, T, P]) Draw(screen S) {
}

// Dispose implements the [Element] interface.
func (e *BaseElement[S, T, P]) Dispose() {
}
