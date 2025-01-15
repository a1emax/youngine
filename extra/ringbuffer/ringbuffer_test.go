package ringbuffer

import (
	"strconv"
	"testing"
)

func TestRingBuffer(t *testing.T) {
	type step struct {
		len        int
		elements   []string
		newElement string // Will not be appended if zero.
	}
	tests := []struct {
		capacity int
		steps    []step
	}{
		{
			capacity: 1,
			steps: []step{
				{len: 0, elements: []string{}, newElement: "a"},
				{len: 1, elements: []string{"a"}, newElement: "b"},
				{len: 1, elements: []string{"b"}},
			},
		},
		{
			capacity: 5,
			steps: []step{
				{len: 0, elements: []string{}, newElement: "a"},
				{len: 1, elements: []string{"a"}, newElement: "b"},
				{len: 2, elements: []string{"a", "b"}, newElement: "c"},
				{len: 3, elements: []string{"a", "b", "c"}, newElement: "d"},
				{len: 4, elements: []string{"a", "b", "c", "d"}, newElement: "e"},
				{len: 5, elements: []string{"a", "b", "c", "d", "e"}, newElement: "f"},
				{len: 5, elements: []string{"b", "c", "d", "e", "f"}, newElement: "g"},
				{len: 5, elements: []string{"c", "d", "e", "f", "g"}, newElement: "h"},
				{len: 5, elements: []string{"d", "e", "f", "g", "h"}, newElement: "i"},
				{len: 5, elements: []string{"e", "f", "g", "h", "i"}, newElement: "j"},
				{len: 5, elements: []string{"f", "g", "h", "i", "j"}, newElement: "k"},
				{len: 5, elements: []string{"g", "h", "i", "j", "k"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.capacity), func(t *testing.T) {
			b := New[string](tt.capacity)
			if b.IsNil() {
				t.Fatalf("new buffer is nil")
			}

			for si, s := range tt.steps {
				if bLen := b.Len(); bLen != s.len {
					t.Fatalf("wrong length on step #%d: %d expected, got %d", si+1, s.len, bLen)
				}
				for i, sElement := range s.elements {
					bElement := b.Get(i)
					if bElement != sElement {
						t.Fatalf("wrong element %d on step #%d: %q expected, got %q", i, si+1, sElement, bElement)
					}
				}
				if s.newElement != "" {
					b.Append(s.newElement)
				}
			}
		})
	}
}
