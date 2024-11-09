package main

import "internal/net"

const (
	NAME    = "hello"
	VERSION = "1.0.0"
	PORT    = "3000"
)

func main() {

	app := net.NewHTTP(NAME, VERSION)
	net.Run(app, PORT)
}
