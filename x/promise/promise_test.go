package promise

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/a1emax/youngine/fault"
)

func TestPromise__Resolve(t *testing.T) {
	expectedSummary, gotSummary := "0+ 1+ 2+ ", ""
	expectedValue := 42

	p, resolve, reject := New[int]()

	n := 3
	resolutionErrs := make([]error, 0, n)
	for i := 0; i < n; i++ {
		i := i

		resolutionErr := fmt.Errorf("resolution error %d", i)
		resolutionErrs = append(resolutionErrs, resolutionErr)

		p.Then(func() {
			gotSummary += strconv.Itoa(i) + "+ "

			panic(resolutionErr)
		}, func(err error) {
			if err == nil {
				t.Fatalf("rejection handler %d was called with nil reason", i)
			} else {
				t.Fatalf("rejection handler %d was called with non-nil reason: %+v", i, err)
			}
		})
	}

	err := fault.Recover(func() {
		_ = p.Value()
	})
	if err == nil {
		t.Fatalf("value is accessible before settlement")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("value getter called before settlement panicked with unexpected error: %+v", err)
	}

	err = fault.Recover(func() {
		resolve(expectedValue)
	})
	if err == nil {
		t.Fatalf("promise was successfully resolved")
	}
	for _, resolutionErr := range resolutionErrs {
		if !errors.Is(err, resolutionErr) {
			t.Fatalf("resolution error is lost: %+v", resolutionErr)
		}
	}

	err = fault.Recover(func() {
		resolve(expectedValue)
	})
	if err == nil {
		t.Fatalf("promise was successfully resolved after it was resolved earlier")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("resolver called after resolution panicked with unexpected error: %+v", err)
	}

	err = fault.Recover(func() {
		reject(nil)
	})
	if err == nil {
		t.Fatalf("promise was successfully rejected after it was resolved earlier")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("rejector called after resolution panicked with unexpected error: %+v", err)
	}

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}

	gotValue := p.Value()
	if gotValue != expectedValue {
		t.Fatalf("wrong value: %d expected, got %d", expectedValue, gotValue)
	}
}

func TestPromise__Reject(t *testing.T) {
	testErr := errors.New("test error")

	expectedSummary, gotSummary := "0- 1- 2- ", ""

	p, resolve, reject := New[int]()

	n := 3
	rejectionErrs := make([]error, 0, n)
	for i := 0; i < n; i++ {
		i := i

		rejectionErr := fmt.Errorf("rejection error %d", i)
		rejectionErrs = append(rejectionErrs, rejectionErr)

		p.Then(func() {
			t.Fatalf("resolution handler %d was called", i)
		}, func(err error) {
			if err == nil {
				t.Fatalf("rejection handler %d was called with nil reason", i)
			}
			if !errors.Is(err, testErr) {
				t.Fatalf("rejection handler %d was called with unexpected reason: %+v", i, err)
			}

			gotSummary += strconv.Itoa(i) + "- "

			panic(rejectionErr)
		})
	}

	err := fault.Recover(func() {
		_ = p.Value()
	})
	if err == nil {
		t.Fatalf("value is accessible before settlement")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("value getter called before settlement panicked with unexpected error: %+v", err)
	}

	err = fault.Recover(func() {
		reject(testErr)
	})
	if err == nil {
		t.Fatalf("promise was successfully rejected")
	}
	for _, rejectionErr := range rejectionErrs {
		if !errors.Is(err, rejectionErr) {
			t.Fatalf("rejection error is lost: %+v", rejectionErr)
		}
	}

	err = fault.Recover(func() {
		reject(testErr)
	})
	if err == nil {
		t.Fatalf("promise was successfully rejected after it was rejected earlier")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("rejector called after rejection panicked with unexpected error: %+v", err)
	}

	err = fault.Recover(func() {
		resolve(0)
	})
	if err == nil {
		t.Fatalf("promise was successfully resolved after it was rejected earlier")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("resolver called after rejection panicked with unexpected error: %+v", err)
	}

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}

	err = fault.Recover(func() {
		_ = p.Value()
	})
	if err == nil {
		t.Fatalf("value is accessible after rejection")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("value getter called after rejection panicked with unexpected error: %+v", err)
	}
}

func TestPromise__Reject__Nil(t *testing.T) {
	expectedSummary, gotSummary := "0- 1- 2- ", ""

	p, resolve, reject := New[int]()

	n := 3
	rejectionErrs := make([]error, 0, n)
	for i := 0; i < n; i++ {
		i := i

		rejectionErr := fmt.Errorf("rejection error %d", i)
		rejectionErrs = append(rejectionErrs, rejectionErr)

		p.Then(func() {
			t.Fatalf("resolution handler %d was called", i)
		}, func(err error) {
			if err == nil {
				t.Fatalf("rejection handler %d was called with nil reason", i)
			}
			if !errors.Is(err, ErrRejected) {
				t.Fatalf("rejection handler %d was called with nil reason with unexpected reason: %+v", i, err)
			}

			gotSummary += strconv.Itoa(i) + "- "

			panic(rejectionErr)
		})
	}

	err := fault.Recover(func() {
		_ = p.Value()
	})
	if err == nil {
		t.Fatalf("value is accessible before settlement")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("value getter called before settlement panicked with unexpected error: %+v", err)
	}

	err = fault.Recover(func() {
		reject(nil)
	})
	if err == nil {
		t.Fatalf("promise was successfully rejected")
	}
	for _, rejectionErr := range rejectionErrs {
		if !errors.Is(err, rejectionErr) {
			t.Fatalf("rejection error is lost: %+v", rejectionErr)
		}
	}

	err = fault.Recover(func() {
		reject(nil)
	})
	if err == nil {
		t.Fatalf("promise was successfully rejected after it was rejected earlier")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("rejector called after rejection panicked with unexpected error: %+v", err)
	}

	err = fault.Recover(func() {
		resolve(0)
	})
	if err == nil {
		t.Fatalf("promise was successfully resolved after it was rejected earlier")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("resolver called after rejection panicked with unexpected error: %+v", err)
	}

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}

	err = fault.Recover(func() {
		_ = p.Value()
	})
	if err == nil {
		t.Fatalf("value is accessible after rejection")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("value getter called after rejection panicked with unexpected error: %+v", err)
	}
}
