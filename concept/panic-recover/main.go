package main

import (
	"fmt"
	"time"
)

func main() {
	f()
	fmt.Println("Return from f() normally.")
}

// Recover is a built-in function that regains control of a panicking goroutine.
// Recover is only useful inside deferred functions. During normal execution,
// a call to recover will return nil and have no other effect. If the current goroutine is panicking,
// a call to recover will capture the value given to panic and resume normal execution.
func f() {
	fmt.Println("calling f()...")
	go func() {
		defer func() {
			// When the goroutine is not panicking, or if the argument supplied to panic was nil,
			// recover returns nil.
			if r := recover(); r != nil {
				fmt.Println("Recover from panic in f(): ", r)
			}
		}()
		fmt.Println("Calling g()...")
		g(0)
		fmt.Println("go func done normally")
	}()

	// go g(0)

	time.Sleep(2 * time.Second)
	fmt.Println("f() call done normally")
}

// Panic is a built-in function that stops the ordinary flow of control and begins panicking.
// When the function F calls panic, execution of F stops, any deferred functions in F are executed normally,
// and then F returns to its caller. To the caller, F then behaves like a call to panic.
// The process continues up the stack until all functions in the current goroutine have returned,
// at which point the program crashes. Panics can be initiated by invoking panic directly.
// They can also be caused by runtime errors, such as out-of-bounds array accesses.
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!!!")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("defer in g(): ", i)
	g(i + 1)
}
