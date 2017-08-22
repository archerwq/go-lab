package main

import "fmt"

func PrintSlice(s []int) {
	fmt.Printf("slice=%v, length=%d, capacity=%d \n", s, len(s), cap(s))
}

func AppendSlice(slice []int, data ...int) []int {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]int, (n+1)*2)
		// The copy function supports copying between slices of different lengths (it will copy only up to the smaller number of elements).
		// In addition, copy can handle source and destination slices that share the same underlying array, handling overlapping slices correctly.
		copy(newSlice, slice)
		slice = newSlice
	}
	// grow the slice
	slice = slice[:n]
	copy(slice[m:n], data)
	return slice
}

func main() {
	// A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).
	// The length is the number of elements referred to by the slice.
	// The capacity is the number of elements in the underlying array (beginning at the element referred to by the slice pointer).
	s := make([]int, 10)
	PrintSlice(s) // [0 0 0 0 0 0 0 0 0 0] 10 10

	// Slicing does not copy the slice's data. It creates a new slice value that points to the original array.
	// This makes slice operations as efficient as manipulating array indices.
	a := s[2:4]
	PrintSlice(a) // [0 0] 2 8

	a = AppendSlice(a, 1, 2, 3)
	PrintSlice(a) // [0 0 1 2 3] 5 8
	PrintSlice(s) // [0 0 0 0 1 2 3 0 0 0] 10 10

	// Modifying the elements (not the slice itself) of a re-slice modifies the elements of the original slice.
	a[1] = 5
	PrintSlice(a) // [0 5 1 2 3] 5 8
	PrintSlice(s) // [0 0 0 5 1 2 3 0 0 0] 10 10

	// We can grow a to its capacity by slicing it again.
	// A slice cannot be grown beyond its capacity. Attempting to do so will cause a runtime panic, just as when indexing outside the bounds of a slice or array.
	// Similarly, slices cannot be re-sliced below zero to access earlier elements in the array.
	a = a[:cap(a)]
	PrintSlice(a) // [0 5 1 2 3 0 0 0] 8 8

	a = AppendSlice(a, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21)
	PrintSlice(a) // [0 5 1 2 3 0 0 0 11 12 13 14 15 16 17 18 19 20 21] 19 40
	PrintSlice(s) // [0 0 0 5 1 2 3 0 0 0] 10 10
}
