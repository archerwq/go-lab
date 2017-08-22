/*
See alse github.com/pkg/errors

Working with Errors in Go 1.13
https://blog.golang.org/go1.13-errors
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	if err := catHosts(); err != nil {
		log.Fatal(err)
	}
}

func catHosts() error {
	data, err := ioutil.ReadFile("/etc/hosts1")
	if err != nil {
		return fmt.Errorf("failed to load hosts: %v", err)
	}
	fmt.Print(string(data))
	return nil
}
