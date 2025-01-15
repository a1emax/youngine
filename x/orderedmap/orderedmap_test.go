package orderedmap

import (
	"testing"
)

func TestOrderedMap(t *testing.T) {
	m := New[string, string]()
	if m.IsNil() {
		t.Fatalf("new map is nil")
	}

	{
		expectedSummary, expectedLen := "", 0

		gotSummary, gotLen := "", m.Len()
		for e := m.First(); !e.IsNil(); e = e.Next() {
			gotSummary += e.Key()
			gotSummary += e.Value()
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong initial summary: %q expected, got %q", expectedSummary, gotSummary)
		}
		if gotLen != expectedLen {
			t.Fatalf("wrong initial length: %d expected, got %d", expectedLen, gotLen)
		}
	}

	m.Set("1", "A")
	m.Set("2", "B")
	m.Set("3", "C")

	{
		expectedSummary, expectedLen := "1A2B3C", 3

		gotSummary, gotLen := "", m.Len()
		for e := m.First(); !e.IsNil(); e = e.Next() {
			gotSummary += e.Key()
			gotSummary += e.Value()
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong summary after inserting: %q expected, got %q", expectedSummary, gotSummary)
		}
		if gotLen != expectedLen {
			t.Fatalf("wrong length after inserting: %d expected, got %d", expectedLen, gotLen)
		}
	}

	m.Set("2", "x")

	{
		key, expectedValue := "2", "x"

		if !m.Contains(key) {
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

		if m.Contains(key) {
			t.Fatalf("deleted key %q found", key)
		}

		gotValue := m.Get(key)
		if gotValue != "" {
			t.Fatalf("wrong value of deleted key %q: expected empty string, got %q", key, gotValue)
		}
	}

	{
		expectedSummary, expectedLen := "2x3C", 2

		gotSummary, gotLen := "", m.Len()
		for e := m.First(); !e.IsNil(); e = e.Next() {
			gotSummary += e.Key()
			gotSummary += e.Value()
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong summary after changing: %q expected, got %q", expectedSummary, gotSummary)
		}
		if gotLen != expectedLen {
			t.Fatalf("wrong length after chaning: %d expected, got %d", expectedLen, gotLen)
		}
	}

	{
		expectedSummary := "3C2x"

		var gotSummary string
		for e := m.Last(); !e.IsNil(); e = e.Prev() {
			gotSummary += e.Key()
			gotSummary += e.Value()
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong reversed summary: %q expected, got %q", expectedSummary, gotSummary)
		}
	}

	{
		expectedSummary, expectedLen := "2x3C", 2

		m := m.ReadOnly()

		gotSummary, gotLen := "", m.Len()
		for e := m.First(); !e.IsNil(); e = e.Next() {
			gotSummary += e.Key()
			gotSummary += e.Value()
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong read-only summary: %q expected, got %q", expectedSummary, gotSummary)
		}
		if gotLen != expectedLen {
			t.Fatalf("wrong read-only length: %d expected, got %d", expectedLen, gotLen)
		}
	}

	{
		expectedSummary, expectedLen := "2x3C", 2

		m := m.Copy()

		gotSummary, gotLen := "", m.Len()
		for e := m.First(); !e.IsNil(); e = e.Next() {
			gotSummary += e.Key()
			gotSummary += e.Value()
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong copy summary: %q expected, got %q", expectedSummary, gotSummary)
		}
		if gotLen != expectedLen {
			t.Fatalf("wrong copy length: %d expected, got %d", expectedLen, gotLen)
		}
	}
}
