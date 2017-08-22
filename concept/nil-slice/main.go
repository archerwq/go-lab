package main

import "fmt"

type File struct {
	ID   string
	Name string
}

type Files []*File

func (fs Files) IDs() []string {
	ids := make([]string, 0)
	for _, f := range fs {
		ids = append(ids, f.ID)
	}
	return ids
}

func main() {
	var files Files
	fmt.Printf("files == nil ? %v\n", files == nil)
	fmt.Println(files)
	fmt.Println(files.IDs()) // 并不会报空指针panic
	files = []*File{
		{"1", "a.txt"},
		{"2", "b.txt"},
	} // 这儿为什么可以这样初始化一个struct指针数组？
	fmt.Println(files.IDs())
}
