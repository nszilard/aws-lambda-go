package pkg

// Fibonacci returns a function that returns successive Fibonacci numbers.
func Fibonacci(digit int) int {
	var a, b = 0, 1

	for i := 0; i < digit; i++ {
		sum := a + b
		a = b
		b = sum
	}

	return a
}
