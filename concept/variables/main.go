package main

import "fmt"

// The var statement declares a list of variables; as in function argument lists, the type is last.
var c, python, java bool

// A var declaration can include initializers, one per variable.
var x, y int = 1, 2

// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
// z := 6

func main() {
	var i int
	// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	var perl, javascript, basic = true, false, true
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	j := 3
	m, n := 4, 5

	fmt.Println(i, c, python, java)
	fmt.Println(x, y, perl, javascript, basic)
	fmt.Println(j, m, n)
}
