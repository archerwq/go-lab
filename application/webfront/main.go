package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/archerwq/go-lab/application/webfront/server"
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
