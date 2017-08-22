package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go parent(ctx)
	time.Sleep(1 * time.Second)
	cancel()

	fmt.Print("type something to exit: ")
	bufio.NewScanner(os.Stdin).Scan()
	fmt.Println("main exit")
}

func parent(ctx context.Context) {
	fmt.Println("parent started")
	done := make(chan struct{})
	go child(ctx, done)
	select {
	case <-done:
		fmt.Println("parent exit normally")
	case <-ctx.Done():
		fmt.Println("parent was canceled")
	}
	fmt.Println("parent exit")
}

func child(ctx context.Context, done chan<- struct{}) {
	fmt.Println("child started")
	time.Sleep(time.Second * 2)
	done <- struct{}{}
	fmt.Println("child exit")
}
