/*
The json package only accesses the exported fields of struct types (those that begin with an uppercase letter).
Therefore only the the exported fields of a struct will be present in the JSON output.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type event struct {
	topic     string
	Version   int                    `json:"version"`
	Data      map[string]interface{} `json:"data"`
	CreatedOn time.Time              `json:"created_on"`
}

func main() {
	e := new(event)
	e.topic = "file_change"
	e.Version = 1
	e.Data = map[string]interface{}{
		"filename": "contacts.xls",
		"fileid":   1111,
		"deleted":  false,
	}
	e.CreatedOn = time.Now()
	fmt.Println(e)

	// marshal
	encoded, err := json.Marshal(e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(encoded))

	// unmarshal
	result := new(event)
	if err := json.Unmarshal(encoded, result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
