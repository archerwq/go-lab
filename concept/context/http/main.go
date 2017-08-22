// https://www.sohamkamani.com/blog/golang/2018-06-17-golang-using-context-cancellation/
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create an HTTP server that listens on port 8000
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// This prints to STDOUT to show that processing has started
		fmt.Fprint(os.Stdout, "processing request\n")

		done := make(chan struct{})
		go calculate(done) // 如果context被cancel掉，这个goroutine会提前退出，因为cancel的时候父goroutine会退出

		// We use `select` to execute a peice of code depending on which
		// channel receives a message first
		select {
		case <-done:
			// If we receive a message after 2 seconds
			// that means the request has been processed
			// We then write this as the response
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			// If the request gets cancelled, log it
			// to STDERR
			fmt.Fprint(os.Stderr, "request cancelled\n")
		}
	}))
}

func calculate(done chan<- struct{}) {
	fmt.Fprint(os.Stderr, "calculating\n")
	<-time.After(5 * time.Second)
	done <- struct{}{}
	fmt.Fprint(os.Stderr, "calculation done\n")
}
