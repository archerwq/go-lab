package main

import (
	"fmt"

	"github.com/archerwq/go-lab/concept/func-options/term"
)

func main() {
	t, err := term.Open("/dev/ttyUSB0")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", t)

	t, err = term.Open("/dev/ttyUSB1", term.RawMode)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", t)

	t, err = term.Open("/dev/ttyUSB2", term.RawMode, term.Speed(16))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", t)
}
