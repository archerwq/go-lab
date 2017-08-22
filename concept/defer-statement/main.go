// When a function returns, its deferred calls are executed in last-in-first-out order.
// A deferred function's arguments are evaluated when the defer statement is evaluated.
// Defer is commonly used to simplify functions that perform various clean-up actions.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// copyFile 用defer机制关闭文件
// 注意在循环中这样做可能会导致系统文件描述符耗尽，因为defer会在最后才执行
// 一种解决方式是循环体里调用另外一个函数处理文件，在新的函数里保证文件被关闭
func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

// slowOperation 利用defer机制记录函数入口和出口以及执行时间
func slowOperation() {
	defer trace("slowOperation")()
	time.Sleep(2 * time.Second)
}

func trace(funcName string) func() {
	start := time.Now()
	log.Printf("enter %s\n", funcName)
	return func() {
		log.Printf("exit %s, (%s)\n", funcName, time.Since(start))
	}
}

// double 利用defer机制记录调用参数和返回值
// defer语句中的函数会在return语句更新返回变量之后再执行，所以可以引用返回值，甚至可以修改返回值
func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	result = x + x
	return
}

func main() {
	copyFile("main.go.swp", "main.go")
	slowOperation()
	double(2)
}
