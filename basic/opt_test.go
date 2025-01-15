package basic

import (
	"testing"
)

func TestOpt__Set(t *testing.T) {
	expected := "test"

	s := SetOpt(expected)

	if !s.IsSet() {
		t.Fatalf("value is not set")
	}

	got := s.Get()
	if got != expected {
		t.Fatalf("%q expected, got %q", expected, got)
	}
}

func TestOpt__NotSet(t *testing.T) {
	var s Opt[string]

	if s.IsSet() {
		t.Fatalf("value is set")
	}

	got := s.Get()
	if got != "" {
		t.Fatalf("zero value expected, got %q", got)
	}
}

func TestOpt_Or__Set(t *testing.T) {
	expected := "test"

	s := SetOpt(expected)

	got := s.Or("default")
	if got != expected {
		t.Fatalf("%q expected, got %q", expected, got)
	}
}

func TestOpt_Or__NotSet(t *testing.T) {
	expected := "test"

	var s Opt[string]

	got := s.Or(expected)
	if got != expected {
		t.Fatalf("%q expected, got %q", expected, got)
	}
}
