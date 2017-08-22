package test

// Average calculate average of an array of integer
func Average(data []float64) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	var sum float64
	for _, f := range data {
		sum += f
	}
	return sum / float64(n)
}

// Fib returns the Nth fib number
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
