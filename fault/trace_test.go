package fault

import (
	"errors"
	"testing"
)

func TestTrace(t *testing.T) {
	type baseError struct {
		error
	}

	baseErr := baseError{errors.New("base error")}

	err := Trace(baseErr)
	if err == nil {
		t.Fatalf("error expected, got nil")
	}

	unwrappedErr := errors.Unwrap(err)
	if unwrappedErr != baseErr {
		t.Fatalf("errors.Unwarap returned unexpected error: %+v", unwrappedErr)
	}

	if !errors.Is(err, baseErr) {
		t.Fatalf("errors.Is returned false")
	}

	var baseTarget baseError
	if !errors.As(err, &baseTarget) {
		t.Fatalf("errors.As returned false")
	}
	if baseTarget != baseErr {
		t.Fatalf("errors.As provided unexpected error: %+v", baseTarget)
	}
}

func TestTrace__Nil(t *testing.T) {
	err := Trace(nil)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}
}
