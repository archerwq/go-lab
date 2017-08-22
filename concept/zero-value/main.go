package main

import "fmt"

type Config struct {
	path  string
	limit int
}

func (c *Config) Path() string {
	// A surprising, but useful, property of nil pointers is you can call methods on types that have a nil value.
	// This can be used to provide default values simply.
	if c == nil {
		return "/usr/home"
	}
	return c.path
}

func main() {
	var b bool
	var i int
	var f float64
	var s string
	var p *int
	var a [10]int
	var sl []int
	var m map[string]int
	var st struct {
		x int
		y int
	}
	var intf interface{}
	var fnc func(int) int

	fmt.Printf("bool: %v\n", b)                                 // false
	fmt.Printf("int: %v\n", i)                                  // 0
	fmt.Printf("float: %v\n", f)                                // 0
	fmt.Printf("string: %v\n", s)                               // ""
	fmt.Printf("zero pointer: == nil? %v\n", p == nil)          // true
	fmt.Printf("array: %v\n", a)                                // [0,0,0,0,0,0,0,0,0,0]
	fmt.Printf("zero slice == nil? %v\n", sl == nil)            // true
	fmt.Printf("slice: %v\n", sl)                               // []
	fmt.Printf("slice len and cap: %v, %v\n", len(sl), cap(sl)) // 0, 0
	fmt.Printf("zero map == nil? %v\n", m == nil)               // true
	fmt.Printf("struct: %v\n", st)                              // {0, 0}
	fmt.Printf("zero interface == nil? %v\n", intf == nil)      // true
	fmt.Printf("zero function == nil? %v\n", fnc == nil)        // true

	var c1 *Config
	var c2 = &Config{
		path: "/export",
	}
	fmt.Println(c1.Path(), c2.Path())
}
