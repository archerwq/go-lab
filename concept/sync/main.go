package main

import (
	"fmt"
	"sync"
	"time"
)

// like CountDownLatch in Java
func waitRoutines() {
	fmt.Println("in waitRoutines")
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("routine done")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("routine done")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("routine done")
	}()

	wg.Wait()
	fmt.Println("exiting waitRoutines")
}

type sig struct{}

// like interrupt in Java
func process(exit chan sig) {
	for {
		select {
		case <-exit:
			fmt.Println("processWithCancel was canceled")
			return
		default:
			fmt.Println("doing something in processWithCancel")
		}
	}
}

func processWithTimeout() {
	exit := make(chan sig)
	go process(exit)
	go process(exit)
	go process(exit)
	// cancel after 1 second
	time.Sleep(time.Second)
	close(exit)
	time.Sleep(time.Second)
	fmt.Println("exit processWithTimeout")
}

func main() {
	waitRoutines()
	processWithTimeout()
}
