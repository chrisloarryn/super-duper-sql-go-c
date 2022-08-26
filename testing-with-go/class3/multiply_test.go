package class3

import "testing"

func TestMultiply(t *testing.T) {
	table := []struct {
		a    int
		b    int
		want int
	}{
		{2, 1, 2},
		{2, 2, 4},
		{2, 3, 6},
		{2, 4, 8},
		{2, 5, 10},
	}

	for _, v := range table {
		got := multiply(v.a, v.b)

		if got != v.want {
			t.Errorf("Expected result (%d), actual result (%d)", v.want, got)
		}
	}

}
