package main

import (
	"flag"
	"fmt"
	"os"

	opt "github.com/mreiferson/go-options"
)

type Options struct {
	HttpAddr string `flag:"http-addr" cfg:"http_addr"`
	KS3Addr  string `flag:"ks3-addr" cfg:"ks3_addr"`
}

func main() {
	options := &Options{}

	flagSet := flag.NewFlagSet("test", flag.ExitOnError)
	flagSet.String("http-addr", options.HttpAddr, "<host>:<port> to listen on for HTTP clients")
	flagSet.String("ks3-addr", options.HttpAddr, "<host>:<port> KS3 address")
	flagSet.Parse(os.Args[1:])

	cfg := map[string]interface{}{"ks3_addr": "store.mbs.cn"}

	opt.Resolve(options, flagSet, cfg)
	fmt.Println(options)
}
