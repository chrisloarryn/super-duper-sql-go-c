package class3

import "testing"

func TestMultiply(t *testing.T) {
	table := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"2x1", 2, 1, 2},
		{"2x2", 2, 2, 4},
		{"2x3", 2, 3, 6},
		{"2x4", 2, 4, 8},
		{"2x5", 2, 5, 10},
	}

	// TODO: executing a single test in the table
	// to execute subtest "go test -v -run 2x3"

	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			got := multiply(v.a, v.b)

			if got != v.want {
				t.Errorf("Expected result (%d), actual result (%d)", v.want, got)
			}
		})
	}

}
