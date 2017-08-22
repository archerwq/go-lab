package main

import "fmt"

func main() {
	// The type *T is a pointer to a T value. Its zero value is nil.
	var p *int
	i, j := 10, 42
	// The & operator generates a pointer to its operand.
	p = &i
	fmt.Println(p) // memory address
	// The * operator denotes the pointer's underlying value.
	fmt.Println(*p) // 10
	*p = 11
	fmt.Println(i) // 11

	p = &j
	*p = *p / 5
	fmt.Println(j) // 8

	// Unlike C, Go has no pointer arithmetic.
}
