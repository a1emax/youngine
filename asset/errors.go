package asset

import (
	"errors"
)

// ErrKindMismatch indicates kind mismatch when loading the same asset.
var ErrKindMismatch = errors.New("asset: kind mismatch")

// ErrKindRebinding indicates attempt to associate the same kind with multiple providers.
var ErrKindRebinding = errors.New("asset: kind rebinding")

// ErrNotLoaded indicates that asset is not loaded for unknown reason (e.g. in case of panic when receiving).
var ErrNotLoaded = errors.New("asset: not loaded")

// ErrTypeMismatch indicates type mismatch when loading typed asset.
var ErrTypeMismatch = errors.New("asset: type mismatch")

// ErrUnknownKind indicates attempt to load asset of unknown kind.
var ErrUnknownKind = errors.New("asset: unknown kind")
