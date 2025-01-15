package fault

import (
	"errors"
)

// ErrIndexOutOfRange indicates attempt to access index that is out of range.
var ErrIndexOutOfRange = errors.New("index out of range")

// ErrIntegerOverflow indicates integer overflow.
var ErrIntegerOverflow = errors.New("integer overflow")

// ErrInvalidArgument indicates attempt to pass invalid argument to function.
var ErrInvalidArgument = errors.New("invalid argument")

// ErrInvalidUse indicates attempt to call function in invalid way.
var ErrInvalidUse = errors.New("invalid use")

// ErrNilPointer indicates attempt to use nil pointer when value is required.
var ErrNilPointer = errors.New("nil pointer")
