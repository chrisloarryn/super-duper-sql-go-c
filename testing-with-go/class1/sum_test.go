package class1_test

import (
	"class1"
	"testing"
)

func TestAdd(t *testing.T) {
	want := 3
	got := class1.Add(1, 2)

	if got != want {
		t.Errorf("Error: want %d, got %d", want, got)
	}
}

func TestAddMultiple(t *testing.T) {
	want := 6
	got := class1.AddMultiple(1, 2, 3)

	if got != want {
		t.Errorf("Error: want %d, got %d", want, got)
	}
}
