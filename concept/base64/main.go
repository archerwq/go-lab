package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	msg := "Hello, 世界"

	// with padding
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(decoded))

	// without padding
	encoded = base64.RawStdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)
	decoded, err = base64.RawStdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(decoded))
}
