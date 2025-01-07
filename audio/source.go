package audio

import (
	"io"
)

// Source of audio data.
type Source interface {
	io.ReadSeeker
}

// FiniteSource is [Source] of known length.
type FiniteSource interface {
	Source

	// Length returns length of data.
	Length() int64
}
