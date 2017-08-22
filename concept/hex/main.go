package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	src := []byte("Hello, 世界")
	encoded := hex.EncodeToString(src)
	fmt.Println(encoded)

	decoded, err := hex.DecodeString(encoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(decoded))
}
