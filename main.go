package main

import (
	"flag"

	server "gitlab.com/hospital-web/server"
)

func main() {
	flag.Parse()

	server.Start()
}
