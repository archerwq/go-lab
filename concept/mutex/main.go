/*
https://mp.weixin.qq.com/s/l315emdX2LayvQtMRMxigA
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	m map[string]int
	// A Mutex is a mutual exclusion lock.
	// The zero value for a RWMutex is an unlocked mutex.
	mux sync.RWMutex
}

func (c *SafeCounter) Inc(key string) {
	// We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock.
	// If the lock is already in use, the calling goroutine blocks until the mutex is available.
	c.mux.Lock()
	c.m[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.RLock()
	// We can use defer to ensure the mutex will be unlocked.
	defer c.mux.RUnlock()
	return c.m[key]
}

func main() {
	counter := SafeCounter{m: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go counter.Inc("key")
	}
	time.Sleep(time.Second)
	fmt.Println(counter.Value("key"))
}
