package main

import "fmt"

type T struct {
	a int
	b float64
	c string
}

type S T

// Override the default print format for a custom type
func (s *S) String() string {
	return fmt.Sprintf("%d/%g/%q", s.a, s.b, s.c)
}

// Another printing technique is to pass a print routine's arguments directly to another such routine.
// The signature of Printf uses the type ...interface{} for its final argument to specify that an arbitrary
// number of parameters (of arbitrary type) can appear after the format.
func log(v ...interface{}) {
	// We write ... after v in the nested call to Println to tell the compiler to treat v as a list of arguments;
	// otherwise it would just pass v as a single slice argument.
	fmt.Println(v...) // a list of arguments
	fmt.Println(v) // a slice
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}

	// When printing a struct, the modified format %+v annotates the fields of the structure with their names,
	// and for any value the alternate format %#v prints the value in full Go syntax.
	fmt.Printf("%v\n", t)  // &{7 -2.35 abc	def}
	fmt.Printf("%+v\n", t) // &{a:7 b:-2.35 c:abc	def}
	fmt.Printf("%#v\n", t) // &main.T{a:7, b:-2.35, c:"abc\tdef"}

	timeZone := map[string]int{"CST": -21600, "PST": -28800, "EST": -18000, "UTC": 0, "MST": -25200}
	fmt.Printf("%v\n", timeZone)  // map[CST:-21600 PST:-28800 EST:-18000 UTC:0 MST:-25200]
	fmt.Printf("%+v\n", timeZone) // map[CST:-21600 PST:-28800 EST:-18000 UTC:0 MST:-25200]
	fmt.Printf("%#v\n", timeZone) // map[string]int{"CST":-21600, "PST":-28800, "EST":-18000, "UTC":0, "MST":-25200}

	// Another handy format is %T, which prints the type of a value.
	fmt.Printf("%T\n", timeZone)

	var s *S = &S{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", s) // 7/-2.35/"abc\tdef"

	log("2018-01-10", "INFO", 1.5, 2, 3)
}
