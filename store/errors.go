package store

import (
	"errors"
)

// ErrIncompatibleType represents attempt to use incompatible data type.
//
// See [IsCompatibleType] for details.
var ErrIncompatibleType = errors.New("store: incompatible type")
