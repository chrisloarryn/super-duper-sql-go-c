package class3

import (
	"reflect"
	"testing"
)

func TestDog(t *testing.T) {
	want := &Dog{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "criollo",
		},
	}

	got := &Dog{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "criollo",
		},
	}

	// TODO: memory pointer
	t.Logf("want %p, got %p", want, got)

	// TODO: could be fixed with reflect, valid in testing, not for CODING

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
