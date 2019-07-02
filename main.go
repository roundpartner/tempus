package main

import (
	"flag"
	"github.com/artyom/autoflags"
	"log"
	"os"
)

var ServiceName = "tempus"

func main() {
	autoflags.Define(&ServerConfig)
	flag.Parse()

	log.SetOutput(os.Stdout)

	ListenAndServe(ServerConfig.Port)
}

var ServerConfig = struct {
	Port int `flag:"port,port number to listen on"`
}{Port: 7373}
