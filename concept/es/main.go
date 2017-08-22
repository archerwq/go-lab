package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(es.Info())
}
