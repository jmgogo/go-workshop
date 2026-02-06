package utils

// private
func add(a int, b int) int {
	return a + b
}

// private
func subtract(a int, b int) int {
	return a - b
}

// public
func AddAndSubtract(a int, b int) (int, int) {
	return add(a, b), subtract(a, b)
}
