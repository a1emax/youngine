package list

import (
	"testing"
)

func TestList(t *testing.T) {
	l := New[string]()
	if l.IsNil() {
		t.Fatalf("new list is nil")
	}

	{
		expectedSummary, expectedLen := "", 0

		gotSummary, gotLen := "", l.Len()
		for e := l.First(); !e.IsNil(); e = e.Next() {
			gotSummary += e.Value()
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong initial summary: %q expected, got %q", expectedSummary, gotSummary)
		}
		if gotLen != expectedLen {
			t.Fatalf("wrong initial length: %d expected, got %d", expectedLen, gotLen)
		}
	}

	l.Prepend("A")              // A
	B := l.Prepend("B")         // BA
	l.Append("C")               // BAC
	D := l.InsertAfter(B, "D")  // BDAC
	E := l.InsertBefore(D, "E") // BEDAC

	{
		expectedSummary, expectedLen := "BEDAC", 5

		gotSummary, gotLen := "", l.Len()
		for e := l.First(); !e.IsNil(); e = e.Next() {
			gotSummary += e.Value()
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong summary after inserting: %q expected, got %q", expectedSummary, gotSummary)
		}
		if gotLen != expectedLen {
			t.Fatalf("wrong length after inserting: %d expected, got %d", expectedLen, gotLen)
		}
	}

	l.Delete(E) // BDAC

	{
		expectedSummary, expectedLen := "BDAC", 4

		gotSummary, gotLen := "", l.Len()
		for e := l.First(); !e.IsNil(); e = e.Next() {
			gotSummary += e.Value()
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong summary after deleting: %q expected, got %q", expectedSummary, gotSummary)
		}
		if gotLen != expectedLen {
			t.Fatalf("wrong length after deleting: %d expected, got %d", expectedLen, gotLen)
		}
	}

	{
		index := 1
		expectedEntry, gotEntry := D, l.Get(index)
		if gotEntry != expectedEntry {
			t.Fatalf("wrong entry with index %d: %q expected, got %q",
				index, expectedEntry.Value(), gotEntry.Value())
		}
	}

	{
		expectedSummary := "CADB"

		var gotSummary string
		for e := l.Last(); !e.IsNil(); e = e.Prev() {
			gotSummary += e.Value()
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong reversed summary: %q expected, got %q", expectedSummary, gotSummary)
		}
	}

	{
		expectedSummary, expectedLen := "BDAC", 4

		l := l.ReadOnly()

		gotSummary, gotLen := "", l.Len()
		for e := l.First(); !e.IsNil(); e = e.Next() {
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
		expectedSummary, expectedLen := "BDAC", 4

		l := l.Copy()

		gotSummary, gotLen := "", l.Len()
		for e := l.First(); !e.IsNil(); e = e.Next() {
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
