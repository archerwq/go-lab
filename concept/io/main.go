package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	read()
	write()
	scan()
}

func read() {
	// read file with ioutil
	data, err := ioutil.ReadFile("main.go")
	check(err)
	fmt.Println(string(data))

	// read file with low-level os package to control over how and what parts of a file are read
	f, err := os.Open("main.go")
	check(err)
	defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes reaeded:\n%s\n", n1, string(b1))

	o2, err := f.Seek(5, 0)
	check(err)
	b2 := make([]byte, 10)
	n2, err := f.Read(b2)
	fmt.Printf("%d bytes @%d reaeded:\n%s\n", n2, o2, string(b2))

	// The io package provides some functions that may be helpful for file reading.
	// For example, reads like the ones below can be more robustly implemented with ReadAtLeast.
	o3, err := f.Seek(5, 0)
	check(err)
	b3 := make([]byte, 10)
	n3, err := io.ReadAtLeast(f, b3, 10)
	check(err)
	fmt.Printf("%d bytes @%d reaeded:\n%s\n", n3, o3, string(b3))

	// The bufio package implements a buffered reader that may be useful both for its efficiency
	// with many small reads and because of the additional reading methods it provides.
	_, err = f.Seek(0, 0)
	check(err)
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes @0 reaeded:\n%s\n", string(b4))
	b5, err := r4.Peek(10)
	check(err)
	fmt.Printf("10 bytes @0 reaeded:\n%s\n", string(b5))
}

func write() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("test1.swp", d1, 0644)
	check(err)

	f, err := os.Create("test2.swp")
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes \n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes \n", n3)
	// Issue a Sync to flush writes to stable storage.
	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	// Use Flush to ensure all buffered operations have been applied to the underlying writer.
	w.Flush()
}

// A line filter is a common type of program that reads input on stdin, processes it,
// and then prints some derived result to stdout. grep and sed are common line filters.
func scan() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.ToUpper(scanner.Text())
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
