package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool = false
	active bool
	hi     string = "Hello Go"
	hello  string

	// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr(?)
	// The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
	// When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
	temperature int    = -10
	age         uint   = 34
	MaxInt      uint64 = 1<<64 - 1
	size        int

	// byte is alias for uint8
	buffer = [...]byte{15, 23, 120}

	// rune is alias for int32, represents a Unicode code point
	chars = []rune("Hi中国!")

	price float64 = 52.16

	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", hi, hi)
	fmt.Printf("Type: %T Value: %v\n", temperature, temperature)
	fmt.Printf("Type: %T Value: %v\n", age, age)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", buffer, buffer)
	fmt.Printf("Type: %T Value: %v\n", chars, chars)
	fmt.Printf("Type: %T Value: %v\n", price, price)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Printf("init values: bool -> [%v], string -> [%v], int -> [%v]\n", active, hello, size)

	ascii := 'a'
	fmt.Printf("%d %[1]c %[1]q %[1]T\n", ascii)
	unicode := '?'
	fmt.Printf("%d %[1]c %[1]q %[1]T\n", unicode)
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q %[1]T\n", newline)
}
