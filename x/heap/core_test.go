package heap

import (
	"testing"
)

type testBase struct {
	a []string
}

func (b *testBase) Len() int {
	return len(b.a)
}

func (b *testBase) Before(i, j int) bool {
	return b.a[i] < b.a[j]
}

func (b *testBase) Swap(i, j int) {
	b.a[i], b.a[j] = b.a[j], b.a[i]
}

func (b *testBase) Append(x string) {
	b.a = append(b.a, x)
}

func (b *testBase) Pop() string {
	n := len(b.a) - 1
	x := b.a[n]
	b.a = b.a[:n]

	return x
}

func TestCore(t *testing.T) {
	c := NewCore(&testBase{})

	checkSummaryAndLen := func(operation, expectedSummary string, expectedLen int) {
		gotSummary, gotLen := "", c.Len()
		for i, n := 0, gotLen; i < n; i++ {
			gotSummary += c.Base.a[i]
		}
		if gotSummary != expectedSummary {
			t.Fatalf("wrong summary after %s: %q expected, got %q", operation, expectedSummary, gotSummary)
		}
		if gotLen != expectedLen {
			t.Fatalf("wrong length after %s: %d expected, got %d", operation, expectedLen, gotLen)
		}
	}

	checkSummaryAndLen("initialization", "", 0)

	c.Insert("B") // B
	c.Insert("C") // BC
	c.Insert("A") // ACB
	c.Insert("D") // ACBD

	checkSummaryAndLen("insertion", "ACBD", 4)

	{
		index := 1
		c.Base.a[index] = "Z"
		c.Update(index)
	}

	checkSummaryAndLen("updating", "ADBZ", 4)

	{
		expectedValue, gotValue := "A", c.MustExtract()
		if gotValue != expectedValue {
			t.Fatalf("wrong extracted value: %q expected, got %q", expectedValue, gotValue)
		}
	}

	checkSummaryAndLen("extraction", "BDZ", 3)

	{
		index := 1
		expectedValue, gotValue := "D", c.Delete(index)
		if gotValue != expectedValue {
			t.Fatalf("wrong deleted value: %q expected, got %q", expectedValue, gotValue)
		}
	}

	checkSummaryAndLen("deletion", "BZ", 2)
}
