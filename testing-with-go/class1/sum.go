package class1

// Add .
func Add(a, b int) int {
	return a + b
}

// AddMultiple .
func AddMultiple(numbers ...int) (sum int) {
	for _, b := range numbers {
		sum += b
	}
	return
}
