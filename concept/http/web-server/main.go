package main

import (
	"net/http"
	_ "net/http/pprof"
	"strings"
)

// SayHello echo the path
func SayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", SayHello)
	if err := http.ListenAndServe(":6060", nil); err != nil {
		panic(err)
	}
}
