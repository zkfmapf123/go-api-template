package main

import "internal/net"

const (
	PORT = "4000"
)

func main() {

	app := net.NewHTTP("core", "1.0.0")
	net.Run(app, PORT)
}