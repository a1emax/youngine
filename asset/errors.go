package asset

import (
	"errors"
)

// ErrKindMismatch represents kind mismatch when loading the same asset.
var ErrKindMismatch = errors.New("asset: kind mismatch")

// ErrKindRebinding represents attempt to associate the same kind with multiple providers.
var ErrKindRebinding = errors.New("asset: kind rebinding")

// ErrTypeMismatch represents type mismatch when loading typed asset.
var ErrTypeMismatch = errors.New("asset: type mismatch")

// ErrUnknownKind represents attempt to load asset of unknown kind.
var ErrUnknownKind = errors.New("asset: unknown kind")
