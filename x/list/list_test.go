package list

import (
	"testing"
)

func TestList(t *testing.T) {
	l := New[string]()
	if l.IsNil() {
		t.Fatalf("new list is nil")
	}

	checkSummaryAndLen := func(l List[string], operation, expectedSummary string, expectedLen int) {
		gotSummary, gotLen := "", l.Len()
		for m := l.First(); !m.IsNil(); m = l.Next(m) {
			gotSummary += l.Get(m)
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong summary after %s: %q expected, got %q", operation, expectedSummary, gotSummary)
		}
		if gotLen != expectedLen {
			t.Fatalf("wrong length after %s: %d expected, got %d", operation, expectedLen, gotLen)
		}
	}

	checkSummaryAndLen(l, "initialization", "", 0)

	l.Prepend("A")              // A
	B := l.Prepend("B")         // BA
	l.Append("C")               // BAC
	D := l.InsertAfter(B, "D")  // BDAC
	E := l.InsertBefore(D, "E") // BEDAC

	checkSummaryAndLen(l, "insertion", "BEDAC", 5)

	{
		expectedValue := "E"
		gotValue := l.Delete(E) // BDAC
		if gotValue != expectedValue {
			t.Fatalf("wrong deleted value: %q expected, got %q", expectedValue, gotValue)
		}
	}

	checkSummaryAndLen(l, "deletion", "BDAC", 4)

	{
		expectedSummary, gotSummary := "CADB", ""
		for m := l.Last(); !m.IsNil(); m = l.Prev(m) {
			gotSummary += l.Get(m)
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong last-prev summary: %q expected, got %q", expectedSummary, gotSummary)
		}
	}

	{
		expectedSummary, gotSummary := "BDAC", ""
		for value := range l.All() {
			gotSummary += value
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong all summary: %q expected, got %q", expectedSummary, gotSummary)
		}
	}

	{
		expectedSummary, gotSummary := "CADB", ""
		for value := range l.Backward() {
			gotSummary += value
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong backward summary: %q expected, got %q", expectedSummary, gotSummary)
		}
	}

	checkSummaryAndLen(l.Copy(), "copying", "BDAC", 4)
}
