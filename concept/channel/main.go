// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
// Like maps and slices, channels must be created before use make(chan T)
// By default, sends and receives block until the other side is ready.
// This allows goroutines to synchronize without explicit locks or condition variables.
// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel: make(chan int, 100)
// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

package main

import (
	"fmt"
	"time"
)

func sum(data []int, c chan int) {
	sum := 0
	for _, n := range data {
		sum += n
	}
	c <- sum
}

func calSum() {
	data := []int{1, 3, 5, -2, 7, 10, 33, 27}
	size := len(data)
	c := make(chan int)
	// Distribute the work between two goroutines.
	// A goroutine is a lightweight thread managed by the Go runtime.
	// go f(x, y) starts a new goroutine running f(x, y)
	// The evaluation of f, x happens in the current goroutine and the execution of f happens in the new goroutine.
	// Goroutines run in the same address space, so access to shared memory must be synchronized.
	go sum(data[:size/2], c)
	go sum(data[size/2:], c)
	// Will block until both goroutines have completed their computation.
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// A sender can close a channel to indicate that no more values will be sent.
	// Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	// Channels aren't like files; you don't usually need to close them.
	// Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	close(c)
}

func fibonacciInf(c, quit chan int) {
	x, y := 0, 1
	for {
		// The select statement lets a goroutine wait on multiple communication operations.
		// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit...")
			return
		}
	}
}

func calFibonacci() {
	c1 := make(chan int, 10)
	fibonacci(10, c1)

	for {
		// Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression.
		// ok is false if there are no more values to receive and the channel is closed.
		x, ok := <-c1
		if !ok {
			break
		}
		fmt.Println(x)
	}

	c2 := make(chan int, 10)
	fibonacci(10, c2)
	// The loop for i := range c receives values from the channel repeatedly until it is closed.
	// This is equivalent to the upper loop, but more concise.
	for i := range c2 {
		fmt.Println(i)
	}

	c3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c3)
		}
		quit <- 0
	}()

	fibonacciInf(c3, quit)
}

func tickAndBoom() {
	tick := time.Tick(1 * time.Second)
	boom := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Print("tick")
		case <-boom:
			fmt.Println("BOOM!!!")
			return
		// The default case in a select is run if no other case is ready.
		default:
			fmt.Print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	calSum()
	calFibonacci()
	tickAndBoom()
}
