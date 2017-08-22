package main

import "fmt"

// A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// A return statement without arguments returns the named return values. This is known as a "naked" return.
	// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
	return
}

func main() {
	a, b := swap("Qiang", "Wang")
	fmt.Println(a, b)

	fmt.Println(split(50))
}
