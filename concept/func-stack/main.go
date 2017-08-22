package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)
}

func f(x int) {
	defer func() {
		fmt.Println(x)
	}()
	fmt.Printf("f(%d)\n", x+0/x) // panic发生时会立即执行本routine的defer语句，然后中断运行，打印panic值和调用栈
	f(x - 1)
}

func printStack() {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false) // print current go routine func stack
	os.Stdout.Write(buf[:n])
}
