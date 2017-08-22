package main

import (
	"encoding/xml"
	"fmt"
)

// <CompleteMultipartUpload>
//   <Part>
//     <PartNumber>1</PartNumber>
//     <ETag>6b8b9a14e4dec5635b8250d29f2df1a5</ETag>
//   </Part>
//    <Part>
//     <PartNumber>2</PartNumber>
//     <ETag>5e67ad5a0811cc979fbf9043df64ced0</ETag>
//   </Part>
// </CompleteMultipartUpload>
type completeMultipartUpload struct {
	XMLName xml.Name `xml:"CompleteMultipartUpload"`
	Parts   []part
}

type part struct {
	XMLName    xml.Name `xml:"Part"`
	PartNumber int64    `xml:"PartNumber"`
	ETag       string   `xml:"ETag"`
}

func getCompleteMultipartUploadBody(totalParts int64) (string, error) {
	parts := make([]part, 0)
	for i := int64(1); i <= totalParts; i++ {
		parts = append(parts, part{PartNumber: i, ETag: "@ETag"})
	}
	v := &completeMultipartUpload{Parts: parts}
	body, err := xml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	fmt.Println("TODO")
}
