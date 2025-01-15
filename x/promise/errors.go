package promise

import (
	"errors"
)

// ErrRejected is default rejection reason.
var ErrRejected = errors.New("promise: rejected")
