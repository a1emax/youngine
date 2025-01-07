package promise

import (
	"errors"
	"testing"

	"github.com/a1emax/youngine/basic"
)

func TestWhen__Resolve__Nil(t *testing.T) {
	expectedSummary, gotSummary := "condition+ callback+ result+ ", ""
	expectedValue := 0

	condition, resolveCondition, _ := New[basic.None]()
	condition.Then(func() {
		gotSummary += "condition+ "
	}, func(err error) {
		if err == nil {
			t.Fatalf("condition rejection handler was called with nil reason")
		} else {
			t.Fatalf("condition rejection handler was called with non-nil reason: %+v", err)
		}
	})

	result := When(condition, func() (Promise[int], error) {
		gotSummary += "callback+ "

		return nil, nil
	})
	result.Then(func() {
		gotValue := result.Value()
		if gotValue != expectedValue {
			t.Fatalf("wrong value: %d expected, got %d", expectedValue, gotValue)
		}

		gotSummary += "result+ "
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		} else {
			t.Fatalf("result rejection handler was called with non-nil reason: %+v", err)
		}
	})

	resolveCondition(basic.None{})
	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestWhen__Resolve__Resolve(t *testing.T) {
	expectedSummary, gotSummary := "condition+ callback+ continuation+ result+ ", ""
	expectedValue := 42

	condition, resolveCondition, _ := New[basic.None]()
	condition.Then(func() {
		gotSummary += "condition+ "
	}, func(err error) {
		if err == nil {
			t.Fatalf("condition rejection handler was called with nil reason")
		} else {
			t.Fatalf("condition rejection handler was called with non-nil reason: %+v", err)
		}
	})

	var resolveContinuation ResolveFunc[int]
	result := When(condition, func() (Promise[int], error) {
		gotSummary += "callback+ "

		var continuation Promise[int]
		continuation, resolveContinuation, _ = New[int]()
		continuation.Then(func() {
			gotSummary += "continuation+ "
		}, func(err error) {
			if err == nil {
				t.Fatalf("continuation rejection handler was called with nil reason")
			} else {
				t.Fatalf("continuation rejection handler was called with non-nil reason: %+v", err)
			}
		})

		return continuation, nil
	})
	result.Then(func() {
		gotValue := result.Value()
		if gotValue != expectedValue {
			t.Fatalf("wrong value: %d expected, got %d", expectedValue, gotValue)
		}

		gotSummary += "result+ "
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		} else {
			t.Fatalf("result rejection handler was called with non-nil reason: %+v", err)
		}
	})

	resolveCondition(basic.None{})
	resolveContinuation(expectedValue)

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestWhen__Resolve__Reject(t *testing.T) {
	continuationErr := errors.New("continuation error")

	expectedSummary, gotSummary := "condition+ callback+ continuation- result- ", ""

	condition, resolveCondition, _ := New[basic.None]()
	condition.Then(func() {
		gotSummary += "condition+ "
	}, func(err error) {
		if err == nil {
			t.Fatalf("condition rejection handler was called with nil reason")
		} else {
			t.Fatalf("condition rejection handler was called with non-nil reason: %+v", err)
		}
	})

	var rejectContinuation RejectFunc
	result := When(condition, func() (Promise[basic.None], error) {
		gotSummary += "callback+ "

		var continuation Promise[basic.None]
		continuation, _, rejectContinuation = New[basic.None]()
		continuation.Then(func() {
			t.Fatalf("continuation resolution handler was called")
		}, func(err error) {
			if err == nil {
				t.Fatalf("continuation rejection handler was called with nil reason")
			}
			if !errors.Is(err, continuationErr) {
				t.Fatalf("continuation rejection handler was called with unexpected reason: %+v", err)
			}

			gotSummary += "continuation- "
		})

		return continuation, nil
	})
	result.Then(func() {
		t.Fatalf("result resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		}
		if !errors.Is(err, continuationErr) {
			t.Fatalf("result rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "result- "
	})

	resolveCondition(basic.None{})
	rejectContinuation(continuationErr)

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestWhen__Resolve__Error(t *testing.T) {
	callbackErr := errors.New("callback error")

	expectedSummary, gotSummary := "condition+ callback+ result- ", ""

	condition, resolveCondition, _ := New[basic.None]()
	condition.Then(func() {
		gotSummary += "condition+ "
	}, func(err error) {
		if err == nil {
			t.Fatalf("condition rejection handler was called with nil reason")
		} else {
			t.Fatalf("condition rejection handler was called with non-nil reason: %+v", err)
		}
	})

	result := When(condition, func() (Promise[basic.None], error) {
		gotSummary += "callback+ "

		return nil, callbackErr
	})
	result.Then(func() {
		t.Fatalf("result resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		}
		if !errors.Is(err, callbackErr) {
			t.Fatalf("result rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "result- "
	})

	resolveCondition(basic.None{})

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestWhen__Reject(t *testing.T) {
	conditionErr := errors.New("condition error")

	expectedSummary, gotSummary := "condition- result- ", ""

	condition, _, rejectCondition := New[basic.None]()
	condition.Then(func() {
		t.Fatalf("condition resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("condition rejection handler was called with nil reason")
		}
		if !errors.Is(err, conditionErr) {
			t.Fatalf("condition rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "condition- "
	})

	result := When(condition, func() (Promise[basic.None], error) {
		t.Fatalf("callback was called")

		return nil, nil
	})
	result.Then(func() {
		t.Fatalf("result resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		}
		if !errors.Is(err, conditionErr) {
			t.Fatalf("result rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "result- "
	})

	rejectCondition(conditionErr)

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestCatch__Resolve(t *testing.T) {
	expectedSummary, gotSummary := "condition+ result- ", ""

	condition, resolveCondition, _ := New[basic.None]()
	condition.Then(func() {
		gotSummary += "condition+ "
	}, func(err error) {
		if err == nil {
			t.Fatalf("condition rejection handler was called with nil reason")
		} else {
			t.Fatalf("condition rejection handler was called with non-nil reason: %+v", err)
		}
	})

	result := Catch(condition, func(err error) (Promise[basic.None], error) {
		if err == nil {
			t.Fatalf("callback was called with nil reason")
		} else {
			t.Fatalf("callback was called with non-nil reason: %+v", err)
		}

		return nil, nil
	})
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

	resolveCondition(basic.None{})

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestCatch__Reject__Nil(t *testing.T) {
	conditionErr := errors.New("condition error")

	expectedSummary, gotSummary := "condition- callback- result+ ", ""
	expectedValue := 0

	condition, _, rejectCondition := New[basic.None]()
	condition.Then(func() {
		t.Fatalf("condition resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("condition rejection handler was called with nil reason")
		}
		if !errors.Is(err, conditionErr) {
			t.Fatalf("condition rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "condition- "
	})

	result := Catch(condition, func(err error) (Promise[int], error) {
		if err == nil {
			t.Fatalf("callback was called with nil reason")
		}
		if !errors.Is(err, conditionErr) {
			t.Fatalf("callback was called with unexpected reason: %+v", err)
		}

		gotSummary += "callback- "

		return nil, nil
	})
	result.Then(func() {
		gotValue := result.Value()
		if gotValue != expectedValue {
			t.Fatalf("wrong value: %d expected, got %d", expectedValue, gotValue)
		}

		gotSummary += "result+ "
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		} else {
			t.Fatalf("result rejection handler was called with non-nil reason: %+v", err)
		}
	})

	rejectCondition(conditionErr)

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestCatch__Reject__Resolve(t *testing.T) {
	conditionErr := errors.New("condition error")

	expectedSummary, gotSummary := "condition- callback- continuation+ result+ ", ""
	expectedValue := 42

	condition, _, rejectCondition := New[basic.None]()
	condition.Then(func() {
		t.Fatalf("condition resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("condition rejection handler was called with nil reason")
		}
		if !errors.Is(err, conditionErr) {
			t.Fatalf("condition rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "condition- "
	})

	var resolveContinuation ResolveFunc[int]
	result := Catch(condition, func(err error) (Promise[int], error) {
		if err == nil {
			t.Fatalf("callback was called with nil reason")
		}
		if !errors.Is(err, conditionErr) {
			t.Fatalf("callback was called with unexpected reason: %+v", err)
		}

		gotSummary += "callback- "

		var continuation Promise[int]
		continuation, resolveContinuation, _ = New[int]()
		continuation.Then(func() {
			gotSummary += "continuation+ "
		}, func(err error) {
			if err == nil {
				t.Fatalf("continuation rejection handler was called with nil reason")
			} else {
				t.Fatalf("continuation rejection handler was called with non-nil reason: %+v", err)
			}
		})

		return continuation, nil
	})
	result.Then(func() {
		gotValue := result.Value()
		if gotValue != expectedValue {
			t.Fatalf("wrong value: %d expected, got %d", expectedValue, gotValue)
		}

		gotSummary += "result+ "
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		} else {
			t.Fatalf("result rejection handler was called with non-nil reason: %+v", err)
		}
	})

	rejectCondition(conditionErr)
	resolveContinuation(expectedValue)

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestCatch__Reject__Reject(t *testing.T) {
	conditionErr := errors.New("condition error")
	continuationErr := errors.New("continuation error")

	expectedSummary, gotSummary := "condition- callback- continuation- result- ", ""

	condition, _, rejectCondition := New[basic.None]()
	condition.Then(func() {
		t.Fatalf("condition resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("condition rejection handler was called with nil reason")
		}
		if !errors.Is(err, conditionErr) {
			t.Fatalf("condition rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "condition- "
	})

	var rejectContinuation RejectFunc
	result := Catch(condition, func(err error) (Promise[basic.None], error) {
		if err == nil {
			t.Fatalf("callback was called with nil reason")
		}
		if !errors.Is(err, conditionErr) {
			t.Fatalf("callback was called with unexpected reason: %+v", err)
		}

		gotSummary += "callback- "

		var continuation Promise[basic.None]
		continuation, _, rejectContinuation = New[basic.None]()
		continuation.Then(func() {
			t.Fatalf("continuation resolution handler was called")
		}, func(err error) {
			if err == nil {
				t.Fatalf("continuation rejection handler was called with nil reason")
			}
			if !errors.Is(err, continuationErr) {
				t.Fatalf("continuation rejection handler was called with unexpected reason: %+v", err)
			}

			gotSummary += "continuation- "
		})

		return continuation, nil
	})
	result.Then(func() {
		t.Fatalf("result resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		}
		if !errors.Is(err, continuationErr) {
			t.Fatalf("result rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "result- "
	})

	rejectCondition(conditionErr)
	rejectContinuation(continuationErr)

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}

func TestCatch__Reject__Error(t *testing.T) {
	conditionErr := errors.New("condition error")
	callbackErr := errors.New("callback error")

	expectedSummary, gotSummary := "condition- callback- result- ", ""

	condition, _, rejectCondition := New[basic.None]()
	condition.Then(func() {
		t.Fatalf("condition resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("condition rejection handler was called with nil reason")
		}
		if !errors.Is(err, conditionErr) {
			t.Fatalf("condition rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "condition- "
	})

	result := Catch(condition, func(err error) (Promise[basic.None], error) {
		if err == nil {
			t.Fatalf("callback was called with nil reason")
		}
		if !errors.Is(err, conditionErr) {
			t.Fatalf("callback was called with unexpected reason: %+v", err)
		}

		gotSummary += "callback- "

		return nil, callbackErr
	})
	result.Then(func() {
		t.Fatalf("result resolution handler was called")
	}, func(err error) {
		if err == nil {
			t.Fatalf("result rejection handler was called with nil reason")
		}
		if !errors.Is(err, callbackErr) {
			t.Fatalf("result rejection handler was called with unexpected reason: %+v", err)
		}

		gotSummary += "result- "
	})

	rejectCondition(conditionErr)

	if gotSummary != expectedSummary {
		t.Fatalf("wrong summary: %q expected, got %q", expectedSummary, gotSummary)
	}
}
