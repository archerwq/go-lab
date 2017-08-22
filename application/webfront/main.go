package main

import (
	"flag"
	"github.com/archerwq/go-lab/applications/webfront/server"
	"log"
	"net/http"
	"time"
)

var (
	httpAddr     = flag.String("http", ":80", "HTTP listen address")
	ruleFile     = flag.String("rules", "", "rule definition file")
	pollInterval = flag.Duration("poll", time.Second*10, "file poll interval")
)

func main() {
	flag.Parse()

	s, err := server.NewServer(*ruleFile, *pollInterval)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(*httpAddr, s)
	if err != nil {
		log.Fatal(err)
	}
}
