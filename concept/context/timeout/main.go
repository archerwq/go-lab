package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Create a new context
	// With a deadline of 100 milliseconds
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Make a request, that will call the baidu homepage
	req, _ := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	// Associate the cancellable context we just created to the request
	req = req.WithContext(ctx)

	// Create a new HTTP client and execute the request
	client := &http.Client{}
	res, err := client.Do(req)
	// If the request failed, log to STDOUT
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	// Print the statuscode if the request succeeds
	fmt.Println("Response received, status code:", res.StatusCode)
}
