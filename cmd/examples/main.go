package main

import (
	"cmd/examples/handlers"
	"internal/net"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "cmd/examples/docs"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

const (
	PORT   = "3000"
	PREFIX = "api"
)

// @title example
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := net.NewHTTP("example", "1.0.0")
	app.Use(logger.New())

	r := app.Group(PREFIX)

	r.Get("/hello", handlers.PingHandlers)

	app.Get("/swagger/*", swagger.HandlerDefault)

	go func() {
		net.Run(app, PORT)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down server...")
	app.Shutdown()
}
