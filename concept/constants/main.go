// Constants in Go are just that—constant. They are created at compile time,
// even when defined as locals in functions, and can only be numbers, characters (runes),
// strings or booleans. Because of the compile-time restriction, the expressions that define
//  them must be constant expressions, evaluatable by the compiler. For instance, 1<<3 is a constant
// expression, while math.Sin(math.Pi/4) is not because the function call to math.Sin needs to happen at run time.
package main

import "fmt"

// Constants are declared like variables, but with the const keyword.
const (
	Pi = 3.14

	// Numeric constants are high-precision values.
	Big = 1 << 100

	Small = Big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

type ByteSize float64

// In Go, enumerated constants are created using the iota enumerator.
// Since iota can be part of an expression and expressions can be implicitly repeated,
// it is easy to build intricate sets of values.
const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // ???
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func main() {
	// Constants cannot be declared using the := syntax.
	const world = "世界"
	fmt.Println("Hello", world)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	// can not compile since constant 1267650600228229401496703205376 overflows int
	// fmt.Println(needInt(Big))

	var b ByteSize = 2017
	fmt.Println(b)

	fmt.Printf("%f %f %f %f %f %f %f %f \n", KB, MB, GB, TB, PB, EB, ZB, YB)
}
