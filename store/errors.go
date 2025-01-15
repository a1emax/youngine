package store

import (
	"errors"
)

// ErrIncompatibleType indicates attempt to use incompatible data type.
//
// See [IsCompatibleType] for details.
var ErrIncompatibleType = errors.New("store: incompatible type")
