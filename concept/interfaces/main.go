// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
// A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
// Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

package main

import (
	"fmt"
	"time"
)

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	// Under the covers, interface values can be thought of as a tuple of a value and a concrete type: (value, type)
	// An interface value holds a value of a specific underlying concrete type.
	var i I

	fmt.Printf("i == nil? %v\n", i == nil) // true
	// Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	// i.M()
	describe(i) // (<nil>, <nil>)

	i = &T{"hello"}
	describe(i) // (&{hello}, *main.T)
	// Calling a method on an interface value executes the method of the same name on its underlying type.
	i.M() // hello

	var t *T
	i = t
	// Note that an interface value that holds a nil concrete value is itself non-nil.
	fmt.Printf("i == nil? %v\n", i == nil) // false
	describe(i)                            // (<nil>, *main.T)
	// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
	i.M() // <nil>

	i = F(128.01)
	describe(i) // (128.01, main.F)
	i.M()       // 128.01

	emptyInterface()
	typeAssertion()

	typeSwitch(1)
	typeSwitch("hello")
	typeSwitch(true)

	stringer()

	customErr()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func emptyInterface() {
	// The interface type that specifies zero methods is known as the empty interface
	// An empty interface may hold values of any type.
	// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
	var k interface{}
	describeAny(k) // (<nil>, <nil>)
	k = 42
	describeAny(k) // (42, int)
	k = "hello"
	describeAny(k) // (hello, string)
}

func describeAny(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func typeAssertion() {
	var k interface{} = "hello"

	// A type assertion provides access to an interface value's underlying concrete value. t := i.(T)
	// If k does not hold string, the statement will trigger a panic.
	s := k.(string)
	fmt.Println(s) // hello

	// A type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
	// If k holds a string, then s will be the underlying value and ok will be true.
	s, ok := k.(string)
	fmt.Println(s, ok) // hello true

	// If not, ok will be false and f will be the zero value of type T, and no panic occurs.
	f, ok := k.(float64)
	fmt.Println(f, ok) // 0 false
}

func typeSwitch(i interface{}) {
	// A type switch is a construct that permits several type assertions in series.
	// A type switch is like a regular switch statement, but the cases in a type switch specify types (not values),
	// and those values are compared against the type of the value held by the given interface value.
	// The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.
	switch v := i.(type) {
	// In each of the int and string cases, the variable v will be of type int or string respectively and hold the value held by i.
	case int:
		fmt.Printf("Twice %v is %v\n", v, 2*v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	// In the default case (where there is no match), the variable v is of the same interface type and value as i.
	default:
		fmt.Printf("I don't know about the type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

// A Stringer is a type that can describe itself as a string.
// The fmt package (and many others) look for this interface to print values.
// type Stringer interface {
//     String() string
// }
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringer() {
	a := Person{"Qiang Wang", 33}
	z := Person{"Johnny Wang", 40}
	fmt.Println(a, z)
}

type MyError struct {
	When time.Time
	What string
}

// The error type is a built-in interface similar to fmt.Stringer:
// type error interface {
//     Error() string
// }
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run(i int) error {
	if i > 0 {
		fmt.Println(i)
		return nil
	}

	// Need to return *MyError here since Error method has pointer receiver.
	return &MyError{
		time.Now(),
		fmt.Sprintf("non positive value: %d", i),
	}
}

func customErr() {
	s := []int{10, -5, 3, 0}
	for _, n := range s {
		// Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
		if err := run(n); err != nil {
			// As with fmt.Stringer, the fmt package looks for the error interface when printing values.
			fmt.Println(err)
		}
	}
}
