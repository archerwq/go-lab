package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// The range form of the for loop iterates over a slice or map (and array, string, channel)
	// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
	// You can skip the index or value by assigning to _.
	// If you only want the index, drop the ", value" entirely. e.g. for i = range pow
	for i, _ := range pow {
		pow[i] = 1 << uint(i)
	}

	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}

	// iterator a map
	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}
