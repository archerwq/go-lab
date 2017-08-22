package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("URLs is missing")
		os.Exit(1)
	}

	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for _ = range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("while fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reding %s: %v", url, err)
		return
	}

	cost := time.Since(start).Nanoseconds() / 1000000
	ch <- fmt.Sprintf("%d\t%d\t%s", cost, nbytes, url)
}
