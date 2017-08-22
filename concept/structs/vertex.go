// Go does not have classes. However, you can define methods on types.

package main

import (
	"fmt"
	"math"
)

// A struct is a collection of fields.
type Vertex struct {
	X float64
	Y float64
}

// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// In this example, the Abs method has a receiver of type Vertex named v.
// You can only declare a method with a receiver whose type is defined in the same package as the method.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Here's Abs written as a regular function with no change in functionality.
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can declare methods with pointer receivers.
// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
// If we remove * from the declaration of Scale, the method will operate on a copy of the original Vertex value.
func (vptr *Vertex) Scale(f float64) {
	vptr.X *= f
	vptr.Y *= f
}

// Here's Scale written as a regular function with no change in functionality.
func Scale(vptr *Vertex, f float64) {
	vptr.X *= f
	vptr.Y *= f
}

func main() {
	var v Vertex
	fmt.Println(v) // {0, 0}, struct's zero value

	v = Vertex{1.0, 2.0}
	// Struct fields are accessed using a dot.
	v.X = 4.1
	fmt.Println(v.X)

	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	// functions with a pointer argument must take a pointer
	// Scale(v, 10) will fail to compile
	Scale(&v, 10)
	fmt.Println(v)
	// while methods with pointer receivers take either a value or a pointer as the receiver when they are called
	// Go interprets the statement v.Scale(10) as (&v).Scale(10) since the Scale method has a pointer receiver.
	v.Scale(10)
	fmt.Println(v)

	p := &v
	// To access the field X of a struct when we have the struct pointer p we could write (*p).X.
	// However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	p.X = 5.3
	fmt.Println(v)

	// Functions that take a value argument must take a value of that specific type
	fmt.Println(Abs(*p))
	// while methods with value receivers take either a value or a pointer as the receiver when they are called
	fmt.Println(p.Abs())
	p.Scale(10)
	fmt.Println(*p)

	v2 := Vertex{X: 3.5}  // Y:0 is implicit
	v3 := Vertex{}        // X:0 and Y:0
	p = &Vertex{1.2, 2.3} // The special prefix & returns a pointer to the struct value.
	fmt.Println(v2, v3, p)
}
