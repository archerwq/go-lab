package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

/*
Command line flag syntax:
	-flag
	-flag=x
	-flag x  // non-boolean flags only
One or two minus signs may be used; they are equivalent.
*/
var user = flag.String("user", "root", "mysql user name")
var host = flag.String("host", "127.0.0.1", "the host mysql resides in")
var port = flag.Int("port", 3306, "the port mysql listening on")
var timeout = flag.Duration("timeout", 30*time.Minute, "session timeout")
var verbose = flag.Bool("verbose", false, "whether show sql")

var passwd string

func init() {
	// Two flags sharing a variable, so we can have a shorthand.
	flag.StringVar(&passwd, "password", "", "password for the user")
	flag.StringVar(&passwd, "p", "", "password for the user (shorthand)")
}

// e.g. flag -user=qwang -p=123456 --host 127.0.0.1 -port 4306 -verbose A B C
// e.g. flag -h
// e.g. flag -help
// e.g. flag --help
func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Args[1:])

	flag.Parse()
	fmt.Println(*user, passwd, *host, *port, *timeout, *verbose)

	// access non flag arguments
	fmt.Println(flag.Args())
}
