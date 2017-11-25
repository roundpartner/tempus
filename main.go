package main

import (
	"github.com/artyom/autoflags"
	"flag"
)

func main() {
	autoflags.Define(&ServerConfig)
	flag.Parse()

	ListenAndServe(ServerConfig.Port)
}

var ServerConfig = struct {
	Port              int    `flag:"port,port number to listen on"`
}{Port: 7373}
