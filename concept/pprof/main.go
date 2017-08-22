/*
When CPU profiling is enabled, the Go program stops about 100 times per second and records
a sample consisting of the program counters on the currently executing goroutine's stack.

https://blog.golang.org/pprof
https://zhuanlan.zhihu.com/p/71529062
https://www.freecodecamp.org/news/how-i-investigated-memory-leaks-in-go-using-pprof-on-a-large-codebase-4bec4325e192/

go tool pprof http://127.0.0.1:8080/debug/pprof/profile -seconds 10
go tool pprof -http=:8081 ~/pprof/pprof.samples.cpu.001.pb.gz

Each box in the graph corresponds to a single function, and the boxes are sized according to
the number of samples in which the function was running. An edge from box X to box Y indicates
that X calls Y; the number along the edge is the number of times that call appears in a sample.
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			localTz()
			doSomething([]byte(`{"a": 1, "b": 2, "c": 3}`))
		}
	}()

	fmt.Println("start api server...")
	panic(http.ListenAndServe(":8080", nil))
}

func doSomething(s []byte) {
	var m map[string]interface{}
	err := json.Unmarshal(s, &m)
	if err != nil {
		panic(err)
	}

	s1 := make([]string, 0)
	s2 := ""
	for i := 0; i < 100; i++ {
		s1 = append(s1, string(s))
		s2 += string(s)
	}
}

func localTz() *time.Location {
	tz, _ := time.LoadLocation("Asia/Shanghai")
	return tz
}
