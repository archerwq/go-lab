package main

import "fmt"

func intSeq() func() int {
	i := 0
	// Go supports anonymous functions, which can form closures.
	// Anonymous functions are useful when you want to define a function inline without having to name it.
	// A closure is a function value that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// We call intSeq, assigning the result (a function) to nextInt.
	// This function value captures its own i value, which will be updated each time we call nextInt.
	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newNextInt := intSeq()
	fmt.Println(newNextInt()) // 1
}
