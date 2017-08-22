package main

import "fmt"

// For any given outer key you must check if the inner map exists, and create it if needed.
func addHit(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}

type Key struct {
	path    string
	country string
}

// Using a single map with a struct key does away with all that complexity.
func addNewHit(m map[Key]int, path, country string) {
	m[Key{path, country}]++
}

func main() {
	hits := make(map[string]map[string]int)
	addHit(hits, "/doc/", "au")
	addHit(hits, "/doc/", "us")
	addHit(hits, "/ref/", "us")
	addHit(hits, "/doc/", "au")
	fmt.Println(hits)
	fmt.Println(hits["/doc/"]["au"])

	newHits := make(map[Key]int)
	addNewHit(newHits, "/doc/", "au")
	addNewHit(newHits, "/doc/", "us")
	addNewHit(newHits, "/ref/", "us")
	addNewHit(newHits, "/doc/", "au")
	fmt.Println(newHits)
	fmt.Println(newHits[Key{"/doc/","au"}])
}
