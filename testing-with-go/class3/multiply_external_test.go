package class3_test

import (
	"class3"
	"testing"
)

func TestMultiplyExported(t *testing.T) {
	want := 6
	got := class3.Multiply(2, 3)

	if got != want {
		t.Errorf("was expected %d, got %d", want, got)
	}
}
