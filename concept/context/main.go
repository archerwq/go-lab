package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func process(ctx context.Context, done chan struct{}) {
	fmt.Println("start processing...")
	// do something
	r := rand.Intn(5)
	fmt.Println(r)
	time.Sleep(time.Duration(r * 1000 * 1000 * 1000))
	fmt.Println("process done")

	done <- struct{}{}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan struct{})
	go process(ctx, done)
	select {
	case <-done:
		fmt.Println("normal exit")
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
