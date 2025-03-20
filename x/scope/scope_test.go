package scope

import (
	"errors"
	"testing"

	"github.com/a1emax/youngine/fault"
)

func TestSetup(t *testing.T) {
	var initSummary, teardownSummary string

	teardown, err := Setup(func(lc Lifecycle) error {
		initSummary += "1"
		lc.Defer(func() {
			teardownSummary += "1"
		})

		initSummary += "2"
		lc.Defer(func() {
			teardownSummary += "2"
		})

		return nil
	})
	if err != nil {
		t.Fatalf("Setup returned unexpected error: %+v", err)
	}

	if expected := "12"; initSummary != expected {
		t.Fatalf("wrong init summary: %q expected, got %q", expected, initSummary)
	}

	if teardownSummary != "" {
		t.Fatalf("empty teardown summary expected, got %q", teardownSummary)
	}

	err = fault.Recover(teardown)
	if err != nil {
		t.Fatalf("TeardownFunc panicked with unexpected error: %+v", err)
	}

	if expected := "21"; teardownSummary != expected {
		t.Fatalf("wrong teardown summary: %q expected, got %q", expected, teardownSummary)
	}
}

func TestSetup__InitError(t *testing.T) {
	testErr := errors.New("test error")

	var initSummary, teardownSummary string

	_, err := Setup(func(lc Lifecycle) error {
		initSummary += "1"
		lc.Defer(func() {
			teardownSummary += "1"
		})

		initSummary += "2"
		lc.Defer(func() {
			teardownSummary += "2"
		})

		return fault.Trace(testErr)
	})
	if err == nil {
		t.Fatalf("Setup did not return error")
	}
	if !errors.Is(err, testErr) {
		t.Fatalf("Setup returned unexpected error: %+v", err)
	}

	if expected := "12"; initSummary != expected {
		t.Fatalf("wrong init summary: %q expected, got %q", expected, initSummary)
	}

	if expected := "21"; teardownSummary != expected {
		t.Fatalf("wrong teardown summary: %q expected, got %q", expected, teardownSummary)
	}
}

func TestSetup__InitPanic(t *testing.T) {
	testErr := errors.New("test error")

	var initSummary, teardownSummary string

	err := fault.Recover(func() {
		_, err := Setup(func(lc Lifecycle) error {
			initSummary += "1"
			lc.Defer(func() {
				teardownSummary += "1"
			})

			initSummary += "2"
			lc.Defer(func() {
				teardownSummary += "2"
			})

			panic(fault.Trace(testErr))
		})
		if err != nil {
			t.Fatalf("Setup returned unexpected error: %+v", err)
		}
	})
	if err == nil {
		t.Fatalf("Setup did not panic")
	}
	if !errors.Is(err, testErr) {
		t.Fatalf("Setup panicked with unexpected error: %+v", err)
	}

	if expected := "12"; initSummary != expected {
		t.Fatalf("wrong init summary: %q expected, got %q", expected, initSummary)
	}

	if expected := "21"; teardownSummary != expected {
		t.Fatalf("wrong teardown summary: %q expected, got %q", expected, teardownSummary)
	}
}

func TestSetup__DeferredFuncPanic(t *testing.T) {
	testErr1 := errors.New("test error 1")
	testErr2 := errors.New("test error 2")

	var initSummary, teardownSummary string

	teardown, err := Setup(func(lc Lifecycle) error {
		initSummary += "1"
		lc.Defer(func() {
			teardownSummary += "1"

			panic(fault.Trace(testErr1))
		})

		initSummary += "2"
		lc.Defer(func() {
			teardownSummary += "2"

			panic(fault.Trace(testErr2))
		})

		return nil
	})
	if err != nil {
		t.Fatalf("Setup returned unexpected error: %+v", err)
	}

	if expected := "12"; initSummary != expected {
		t.Fatalf("wrong init summary: %q expected, got %q", expected, initSummary)
	}

	if teardownSummary != "" {
		t.Fatalf("empty teardown summary expected, got %q", teardownSummary)
	}

	err = fault.Recover(teardown)
	if err == nil {
		t.Fatalf("TeardownFunc did not panic")
	}
	if !errors.Is(err, testErr1) {
		t.Fatalf("%+v is lost: %+v", testErr1, err)
	}
	if !errors.Is(err, testErr2) {
		t.Fatalf("%+v is lost: %+v", testErr2, err)
	}

	if expected := "21"; teardownSummary != expected {
		t.Fatalf("wrong teardown summary: %q expected, got %q", expected, teardownSummary)
	}
}

func TestSetup__PanicEverywhere(t *testing.T) {
	initErr := errors.New("init error")
	testErr1 := errors.New("test error 1")
	testErr2 := errors.New("test error 2")

	var initSummary, teardownSummary string

	err := fault.Recover(func() {
		_, err := Setup(func(lc Lifecycle) error {
			initSummary += "1"
			lc.Defer(func() {
				teardownSummary += "1"

				panic(fault.Trace(testErr1))
			})

			initSummary += "2"
			lc.Defer(func() {
				teardownSummary += "2"

				panic(fault.Trace(testErr2))
			})

			panic(fault.Trace(initErr))
		})
		if err != nil {
			t.Fatalf("Setup returned unexpected error: %+v", err)
		}
	})
	if err == nil {
		t.Fatalf("Setup did not panic")
	}
	if !errors.Is(err, initErr) {
		t.Fatalf("%+v is lost: %+v", initErr, err)
	}
	if !errors.Is(err, testErr1) {
		t.Fatalf("%+v is lost: %+v", testErr1, err)
	}
	if !errors.Is(err, testErr2) {
		t.Fatalf("%+v is lost: %+v", testErr2, err)
	}

	if expected := "12"; initSummary != expected {
		t.Fatalf("wrong init summary: %q expected, got %q", expected, initSummary)
	}

	if expected := "21"; teardownSummary != expected {
		t.Fatalf("wrong teardown summary: %q expected, got %q", expected, teardownSummary)
	}
}

func TestSetup__DeferAfterSetup(t *testing.T) {
	var _lc Lifecycle
	teardown, err := Setup(func(lc Lifecycle) error {
		_lc = lc

		return nil
	})
	if err != nil {
		t.Fatalf("Setup returned unexpected error: %+v", err)
	}

	err = fault.Recover(func() {
		_lc.Defer(func() {})
	})
	if err == nil {
		t.Fatalf("Lifecycle.Defer was successfully called after initialization")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("Lifecycle.Defer called after initialization panicked with unexpected error: %+v", err)
	}

	teardown()
}

func TestSetup__DeferInDispose(t *testing.T) {
	var n int
	teardown, err := Setup(func(lc Lifecycle) error {
		lc.Defer(func() {
			lc.Defer(func() {
				n++
			})
		})

		return nil
	})
	if err != nil {
		t.Fatalf("Setup returned unexpected error: %+v", err)
	}

	err = fault.Recover(teardown)
	if err == nil {
		t.Fatalf("TeardownFunc did not panic")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("TeardownFunc panicked with unexpected error: %+v", err)
	}

	if n > 0 {
		t.Fatalf("Lifecycle.Defer was successfully called in TeardownFunc")
	}
}

func TestSetup__MultipleDispose(t *testing.T) {
	var n int
	teardown, err := Setup(func(lc Lifecycle) error {
		lc.Defer(func() {
			n++
		})

		return nil
	})
	if err != nil {
		t.Fatalf("Setup returned unexpected error: %+v", err)
	}

	if n > 0 {
		t.Fatalf("deferred function was called too early")
	}

	teardown()
	err = fault.Recover(teardown)
	if err == nil {
		t.Fatalf("TeardownFunc did not panic on second call")
	}
	if !errors.Is(err, fault.ErrInvalidUse) {
		t.Fatalf("TeardownFunc panicked with unexpected error: %+v", err)
	}

	if n > 1 {
		t.Fatalf("deferred function was called more than once")
	}
}

func TestMustSetup(t *testing.T) {
	var initSummary, teardownSummary string

	var teardown TeardownFunc
	err := fault.Recover(func() {
		teardown = MustSetup(func(lc Lifecycle) {
			initSummary += "1"
			lc.Defer(func() {
				teardownSummary += "1"
			})

			initSummary += "2"
			lc.Defer(func() {
				teardownSummary += "2"
			})
		})
	})
	if err != nil {
		t.Fatalf("MustSetup panicked with unexpected error: %+v", err)
	}

	if expected := "12"; initSummary != expected {
		t.Fatalf("wrong init summary: %q expected, got %q", expected, initSummary)
	}

	if teardownSummary != "" {
		t.Fatalf("empty teardown summary expected, got %q", teardownSummary)
	}

	err = fault.Recover(teardown)
	if err != nil {
		t.Fatalf("TeardownFunc panicked with unexpected error: %+v", err)
	}

	if expected := "21"; teardownSummary != expected {
		t.Fatalf("wrong teardown summary: %q expected, got %q", expected, teardownSummary)
	}
}
