package main

import "fmt"

// In Go the line between inheritance and composition is pretty blurry in comparison with Java.
// There is no extends keyword. Syntactically, inheritance looks almost identical to composition.
// The only difference between composition and inheritance in Go, is a struct which inherits from
// another struct can directly access the methods and fields of the parent struct.

type Pet struct {
	name string
}

func (p *Pet) Speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

func (p *Pet) Name() string {
	return p.name
}

type Dog struct {
	Pet
	Breed string
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and I am a %v", d.Pet.Speak(), d.Breed)
}

func main() {
	d := Dog{Pet: Pet{name: "spot"}, Breed: "pointer"}
	fmt.Println(d.Name())
	fmt.Println(d.Speak())
}
