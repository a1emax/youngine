package promise

import (
	"errors"
	"strconv"
	"testing"

	"github.com/a1emax/youngine/basic"
)

func TestAll__Resolve(t *testing.T) {
	expectedSummary, gotSummary := "item0+ item1+ item2+ result+ ", ""

	n := 3
	items := make([]Thener, n)
	itemSettlers := make([]struct {
		resolve ResolveFunc[basic.None]
		reject  RejectFunc
	}, n)
	for i := 0; i < n; i++ {
		i := i

		items[i], itemSettlers[i].resolve, itemSettlers[i].reject = New[basic.None]()
		items[i].Then(func() {
			gotSummary += "item" + strconv.Itoa(i) + "+ "
		}, func(err error) {
			if err == nil {
				t.Fatalf("item%d rejection handler was called with nil reason", i)
			} else {
				t.Fatalf("item%d rejection handler was called with non-nil reason: %+v", i, err)
			}
		})
	}

	result := All(items...)
	result.Then(func() {
		gotSummary += "result+ "

		_ = result.Value()
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		} else {
			t.Fatalf("result rejection handler was called with non-nil reason: %+v", err)
		}
	})

	for _, itemSettler := range itemSettlers {
		itemSettler.resolve(basic.None{})
	}

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestAll__RejectFirst(t *testing.T) {
	testErr := errors.New("test error")

	expectedSummary, gotSummary := "item0- result- item1+ item2+ ", ""

	n := 3
	items := make([]Thener, n)
	itemSettlers := make([]struct {
		resolve ResolveFunc[basic.None]
		reject  RejectFunc
	}, n)
	for i := 0; i < n; i++ {
		i := i

		items[i], itemSettlers[i].resolve, itemSettlers[i].reject = New[basic.None]()
		items[i].Then(func() {
			if i == 0 {
				t.Fatalf("item%d resoulution handler was called", i)
			} else {
				gotSummary += "item" + strconv.Itoa(i) + "+ "
			}
		}, func(err error) {
			if i == 0 {
				if err == nil {
					t.Fatalf("item%d rejection handler was called with nil reason", i)
				}
				if !errors.Is(err, testErr) {
					t.Fatalf("item%d rejection handler was called with unexpected reason: %+v", i, err)
				}

				gotSummary += "item" + strconv.Itoa(i) + "- "
			} else {
				if err == nil {
					t.Fatalf("item%d rejection handler was called with nil reason", i)
				} else {
					t.Fatalf("item%d rejection handler was called with non-nil reason: %+v", i, err)
				}
			}
		})
	}

	result := All(items...)
	result.Then(func() {
		t.Fatalf("result resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		}
		if !errors.Is(err, ErrRejected) {
			t.Fatalf("result rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "result- "
	})

	for i, itemSettler := range itemSettlers {
		if i == 0 {
			itemSettler.reject(testErr)
		} else {
			itemSettler.resolve(basic.None{})
		}
	}

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestAny__Resolve(t *testing.T) {
	expectedSummary, gotSummary := "item0+ result+ item1+ item2+ ", ""
	expectedIndex := 0

	n := 3
	items := make([]Thener, n)
	itemSettlers := make([]struct {
		resolve ResolveFunc[basic.None]
		reject  RejectFunc
	}, n)
	for i := 0; i < n; i++ {
		i := i

		items[i], itemSettlers[i].resolve, itemSettlers[i].reject = New[basic.None]()
		items[i].Then(func() {
			gotSummary += "item" + strconv.Itoa(i) + "+ "
		}, func(err error) {
			if err == nil {
				t.Fatalf("item%d rejection handler was called with nil reason", i)
			} else {
				t.Fatalf("item%d rejection handler was called with non-nil reason: %+v", i, err)
			}
		})
	}

	result := Any(items...)
	result.Then(func() {
		gotSummary += "result+ "

		gotIndex := result.Value()
		if gotIndex != expectedIndex {
			t.Fatalf("wrong index: %d expected, got %d", expectedIndex, gotIndex)
		}
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		} else {
			t.Fatalf("result rejection handler was called with non-nil reason: %+v", err)
		}
	})

	for _, f := range itemSettlers {
		f.resolve(basic.None{})
	}

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestAny__RejectFirst(t *testing.T) {
	testErr := errors.New("test error")

	expectedSummary, gotSummary := "item0- item1+ result+ item2+ ", ""
	expectedIndex := 1

	n := 3
	items := make([]Thener, n)
	itemSettlers := make([]struct {
		resolve ResolveFunc[basic.None]
		reject  RejectFunc
	}, n)
	for i := 0; i < n; i++ {
		i := i

		items[i], itemSettlers[i].resolve, itemSettlers[i].reject = New[basic.None]()
		items[i].Then(func() {
			if i == 0 {
				t.Fatalf("item%d resolution handler was called", i)
			} else {
				gotSummary += "item" + strconv.Itoa(i) + "+ "
			}
		}, func(err error) {
			if i == 0 {
				if err == nil {
					t.Fatalf("item%d rejection handler was called with nil reason", i)
				}
				if !errors.Is(err, testErr) {
					t.Fatalf("item%d rejection handler was called with unexpected reason: %+v", i, err)
				}

				gotSummary += "item" + strconv.Itoa(i) + "- "
			} else {
				if err == nil {
					t.Fatalf("item%d rejection handler was called with nil reason", i)
				} else {
					t.Fatalf("item%d rejection handler was called with non-nil reason: %+v", i, err)
				}
			}
		})
	}

	result := Any(items...)
	result.Then(func() {
		gotSummary += "result+ "

		gotIndex := result.Value()
		if gotIndex != expectedIndex {
			t.Fatalf("wrong index: %d expected, got %d", expectedIndex, gotIndex)
		}
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		} else {
			t.Fatalf("result rejection handler was called with non-nil reason: %+v", err)
		}
	})

	for i, itemSettler := range itemSettlers {
		if i == 0 {
			itemSettler.reject(testErr)
		} else {
			itemSettler.resolve(basic.None{})
		}
	}

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestAny__RejectAll(t *testing.T) {
	testErr := errors.New("test error")

	expectedSummary, gotSummary := "item0- item1- item2- result- ", ""

	n := 3
	items := make([]Thener, n)
	itemSettlers := make([]struct {
		resolve ResolveFunc[basic.None]
		reject  RejectFunc
	}, n)
	for i := 0; i < n; i++ {
		i := i

		items[i], itemSettlers[i].resolve, itemSettlers[i].reject = New[basic.None]()
		items[i].Then(func() {
			t.Fatalf("item%d resolution handler was called", i)
		}, func(err error) {
			if err == nil {
				t.Fatalf("item%d rejection handler was called with nil reason", i)
			}
			if !errors.Is(err, testErr) {
				t.Fatalf("item%d rejection handler was called with unexpected reason: %+v", i, err)
			}

			gotSummary += "item" + strconv.Itoa(i) + "- "
		})
	}

	result := Any(items...)
	result.Then(func() {
		t.Fatalf("result resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		}
		if !errors.Is(err, ErrRejected) {
			t.Fatalf("result rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "result- "
	})

	for _, itemSettler := range itemSettlers {
		itemSettler.reject(testErr)
	}

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}
