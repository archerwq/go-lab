package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	// Whenever one struct field is embedded in another,
	// Go gives us the option to access the embedded fields
	// as if they were part of the outer struct.
	// fmt.Println("Author: ", p.author.fullName())
	// fmt.Println("Bio: ", p.author.bio)
	fmt.Println("Author: ", p.fullName())
	fmt.Println("Bio: ", p.bio)
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post{
		title:   "Inheritance in Go",
		content: "Go supports composition instead of inheritance",
		author:  author1,
	}
	post1.details()
}
