package class2

import "testing"

func TestFibonacciFor(t *testing.T) {
	t.Run("fibonacciFor(10) should return 55", func(t *testing.T) {
		want := 55
		got := fibonacciFor(10)
		if got != want {
			t.Errorf("fibonacciFor(10) = %d, want %d", got, want)
		}
	})
}

func TestFibonacciRec(t *testing.T) {
	t.Run("fibonacciRec(10) should return 55", func(t *testing.T) {
		want := 55
		got := fibonacciRec(10)
		if got != want {
			t.Errorf("fibonacciRec(10) = %d, want %d", got, want)
		}
	})
}

func TestFibonnaciRecMem(t *testing.T) {
	t.Run("fibonacciRecMem(10) should return 55", func(t *testing.T) {
		want := 55
		got := fibonacciRecMem(10)
		if got != want {
			t.Errorf("fibonacciRecMem(10) = %d, want %d", got, want)
		}
	})
}

func BenchmarkFibonacciFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciFor(30)
	}
}

func BenchmarkFibonacciRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRec(30)
	}
}

func BenchmarkFibonacciRecMem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecMem(30)
	}
}
