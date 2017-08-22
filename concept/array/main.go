package main

import "fmt"

func main() {
	var a [2]string
	fmt.Println(len(a), a) // 2, []
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
