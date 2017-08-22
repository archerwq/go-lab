// Map keys may be of any type that is comparable. The language spec defines this precisely, 
// but in short, comparable types are boolean, numeric, string, pointer, channel, and interface types, 
// and structs or arrays that contain only those types. Notably absent from the list are slices, maps, and functions; 
// these types cannot be compared using ==, and may not be used as map keys.

// Maps are not safe for concurrent use: it's not defined what happens when you read and write to them simultaneously. 
// If you need to read from and write to a map from concurrently executing goroutines, the accesses must be mediated by 
// some kind of synchronization mechanism. One common way to protect maps is with sync.RWMutex.

// More on map: https://blog.golang.org/go-maps-in-action

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
	var m map[string]Vertex
	fmt.Println(m == nil) // true

	// The make function returns a map of the given type, initialized and ready for use.
	m = make(map[string]Vertex)
	fmt.Println(m) // map[]

	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967}]
	fmt.Println(m["Bell Labs"]) // {40.68433 -74.39967}

	m["Bell Labs"] = Vertex{50.68433, -84.39967}
	fmt.Println(m)

	// Map literals are like struct literals, but the keys are required.
	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	m2 := map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	// insert or update an element in map m2
	m2["Zuora"] = Vertex{50.62341, -121.12345}
	fmt.Println(m2)

	// retrieve an element
	zuoraVertex := m2["Zuora"]
	fmt.Printf("Zuora: %v \n", zuoraVertex)

	// builtin delete function to delete an element of a map
	delete(m2, "Zuora")

	// return Vertex{0, 0} !!!
	// Zero value of the value type is returned if key is absent.
	zuoraVertex = m2["Zuora"]
        fmt.Printf("Zuora: %v \n", zuoraVertex)

	// two-value version to test a key is present
	elem, ok := m2["Zuora"]
	fmt.Println("The value:", elem, "Present?", ok)
}
