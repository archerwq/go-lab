package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 5
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)

	// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax),
	// the variable's type is inferred from the value on the right hand side.
	var i int32 = 34
	j := i
	fmt.Printf("j is of type %T\n", j)

	// But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or complex128
	// depending on the precision of the constant
	v := 42
	fmt.Printf("v is of type %T\n", v)
	var pi = 3.142
	fmt.Printf("pi is of type %T\n", pi)

	// type switch
	checkType(x)  // integer 3
	checkType(&x) // pointer to integer 3

	// A type assertion takes an interface value and extracts from it a value of the specified explicit type. 
	// The syntax borrows from the clause opening a type switch, but with an explicit type rather than the type keyword
	var m interface{} = 3
	v1, ok1 := m.(int)
	fmt.Println(v1, ok1) // 3 true
	v2, ok2 := m.(bool)
	// If the type assertion fails, v2 will still exist and be of type bool, but it will have the zero value, false.
	fmt.Println(v2, ok2) // false false
}

func checkType(v interface{}) {
	// Type switches are a form of conversion: they take an interface and, 
	// for each case in the switch, in a sense convert it to the type of that case. 
	switch t := v.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t)
	}
}
