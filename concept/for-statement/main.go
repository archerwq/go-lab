package main

import "fmt"

func sum1() int {
	sum := 0
	// Go has only one looping construct, the for loop.
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum
}

func sum2() int {
	sum := 1
	// like 'while' in other language, equals 'for ; sum < 1000; {'
	for sum < 1000 {
		sum += sum
	}
	return sum
}

func sum3() int {
	sum := 1
	// loops forever
	for {
		sum += sum
		if sum > 1000 {
			break
		}
	}
	return sum
}

func main() {
	fmt.Println(sum1())
	fmt.Println(sum2())
	fmt.Println(sum3())
}
