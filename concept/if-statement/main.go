package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// like for loops, the expression need not be surrounded by parentheses ( )
	//  but the braces { } are required
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, limit float64) float64 {
	// if statement can start with a short statement to execute before the condition.
	// Variables declared by the statement are only in scope until the end of the if.
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		// Variables declared inside an if short statement are also available inside any of the else blocks.
		fmt.Printf("%g >= %g\n", v, limit)
	}
	// can not use v here
	return limit
}

func main() {
	fmt.Println(sqrt(8))
	fmt.Println(sqrt(-64))
	fmt.Println(pow(2, 3, 10))
	fmt.Println(pow(3, 3, 20))
}
