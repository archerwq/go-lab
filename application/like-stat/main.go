package main

import "fmt"

type Person struct {
	Name  string
	Likes []string
}

func buildLikes(people []*Person) map[string][]*Person {
	likes := make(map[string][]*Person)
	// It's safe we don't check if people is nil
	for _, p := range people {
		for _, l := range p.Likes {
			// Appending to a nil slice just allocates a new slice.
			likes[l] = append(likes[l], p)
		}
	}
	return likes
}

func main() {
	var people []*Person
	fmt.Println(buildLikes(people))

	people = []*Person{
		&Person{"Qiang Wang", []string{"Basketball", "Travel", "PES"}},
		&Person{"San Zhang", []string{"Basketball", "Football", "Food", "Programming"}},
		&Person{"Si Li", []string{"Food", "Shopping", "Walking"}},
	}

	fmt.Println(buildLikes(people))
}
