// Executable commands must always use package main.
package main

import (
	"fmt"
	"math"

	"github.com/archerwq/go-lab/lib/strutil"
)

// One package can only have one main function.
func main() {
	fmt.Printf("Hello Go!\n")
	fmt.Printf("%s\n", strutil.Reverse("Hello Go!"))

	fmt.Printf("Pi value is %f\n", math.Pi)
	fmt.Printf("Thank you!\n")
}
