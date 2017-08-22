package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("enter main...")
	defer func() {
		fmt.Println("deferred func call in main")
		// if panic happens in sub routine, main routine will exit as well, without defer execution.
		// so will not recover here!!! it can only recover in that sub routine.
		if r := recover(); r != nil {
			fmt.Println("revovering the panic: ", r)
		}
	}()
	c := make(chan int)
	// this routine will exit immediately without defer execution if main routine exits.
	// we need to use channel to make sure main will wait until this routine done
	go f(c)
	// go safelyDo(c)
	<-c
	fmt.Println("exit main")
}

func f(c chan int) {
	fmt.Println("enter f...")
	defer fmt.Println("deferred func call in f")
	func() {
		fmt.Println("enter goroutine...")
		defer func() {
			fmt.Println("deferred func call in goroutine")
			c <- 1
		}()
		for i := 1; i < 6; i++ {
			if i == 5 {
				// When panic is called, including implicitly for run-time errors such as
				// indexing a slice out of bounds or failing a type assertion, it immediately
				// stops execution of the current function and begins unwinding the stack of
				// the goroutine, running any deferred functions along the way. If that unwinding
				// reaches the top of the goroutine's stack, the program dies.
				panic("i == 5")
			}
			fmt.Println(i)
			time.Sleep(time.Second)
		}
		fmt.Println("exit goroutine")
	}()
	fmt.Println("exit f")
}

func safelyDo(c chan int) {
	// one application of recover is to shut down a failing goroutine inside a
	//  without killing the other executing goroutines.
	defer func() {
		fmt.Println("deferred func call in safelyDo")
		// A call to recover stops the unwinding and returns the argument passed to panic.
		// Because the only code that runs while unwinding is inside deferred functions,
		// recover is only useful inside deferred functions.
		if r := recover(); r != nil {
			fmt.Println("work failed: ", r)
		}
	}()
	f(c)
}
