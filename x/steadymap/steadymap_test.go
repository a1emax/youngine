package steadymap

import (
	"testing"
)

func TestSteadyMap(t *testing.T) {
	m := New[string, string]()
	if m.IsNil() {
		t.Fatalf("new map is nil")
	}

	checkSummaryAndLen := func(m SteadyMap[string, string], operation, expectedSummary string, expectedLen int) {
		gotSummary, gotLen := "", m.Len()
		for key, value := range m.All() {
			gotSummary += key
			gotSummary += value
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong summary after %s: %q expected, got %q", operation, expectedSummary, gotSummary)
		}
		if gotLen != expectedLen {
			t.Fatalf("wrong length after %s: %d expected, got %d", operation, expectedLen, gotLen)
		}
	}

	checkSummaryAndLen(m, "initialization", "", 0)

	m.Set("1", "A")
	m.Set("2", "B")
	m.Set("3", "C")

	checkSummaryAndLen(m, "insertion", "1A2B3C", 3)

	m.Set("2", "x")

	{
		key, expectedValue := "2", "x"

		if !m.Has(key) {
			t.Fatalf("changed key %q not found", key)
		}

		gotValue := m.Get(key)
		if gotValue != expectedValue {
			t.Fatalf("wrong value of changed key %q: %q expected, got %q", key, expectedValue, gotValue)
		}
	}

	m.Delete("1")

	{
		key := "1"

		if m.Has(key) {
			t.Fatalf("deleted key %q found", key)
		}

		gotValue := m.Get(key)
		if gotValue != "" {
			t.Fatalf("wrong value of deleted key %q: expected empty string, got %q", key, gotValue)
		}
	}

	checkSummaryAndLen(m, "changing", "2x3C", 2)

	{
		expectedSummary, gotSummary := "3C2x", ""
		for key, value := range m.Backward() {
			gotSummary += key
			gotSummary += value
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong backward summary: %q expected, got %q", expectedSummary, gotSummary)
		}
	}

	checkSummaryAndLen(m.Copy(), "copying", "2x3C", 2)
}
