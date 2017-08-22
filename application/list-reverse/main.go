package main

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

func InitList(values []interface{}) *Node {
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

func PrintList(list *Node) {
	for head := list; head != nil; head = head.next {
		fmt.Print(head.value, " -> ")
	}
	fmt.Print("nil")
	fmt.Println()
}

func ReverseList(list *Node) *Node {
	head := list
	if head.next == nil {
		return head
	}
	var p1 *Node = nil
	var p2 *Node = head
	for p2 != nil {
		p3 := p2.next
		p2.next = p1
		p1 = p2
		p2 = p3
	}
	return p1
}

func main() {
	list := InitList([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	PrintList(list)
	reversed := ReverseList(list)
	PrintList(reversed)
}
