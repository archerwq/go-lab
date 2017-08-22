package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

func InitLink(values []interface{}) *Node {
	var head, tail *Node
	for i := len(values) - 1; i >= 0; i-- {
		head = &Node{
			values[i],
			tail,
		}
		tail = head
	}
	return head
}

func PrintLink(link *Node) {
	for p := link; p.next != nil; p = p.next {
		fmt.Printf("%v -> ", p.value)
	}
	fmt.Println("nil")
}

func HasCircle(link *Node) bool {
	visited := make(map[*Node]bool)
	for p := link; p != nil; p = p.next {
		// if a key is absent in a map, zero value (false here) will be returned
		if visited[p] {
			return true
		}
		visited[p] = true
	}
	return false
}

func main() {
	link := InitLink([]interface{}{1, 2, 3, 4, 5, 6, 7})
	PrintLink(link)
	fmt.Println(HasCircle(link))
}
