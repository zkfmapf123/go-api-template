package net

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	PING_PATH    = "/ping"
	VERSION_PATH = "/version"
)

type HTTPConnector struct {
	App *fiber.App
}

func NewHTTP(name, version string) *fiber.App {

	app := fiber.New(
		fiber.Config{
			Prefork:       false,
			CaseSensitive: true,
			StrictRouting: true,
			ServerHeader:  name,
			AppName:       fmt.Sprintf("%s v%s", name, version),
		},
	)

	setPing(app, "success")
	setVersion(app, version)
	return app
}

func setPing(app *fiber.App, message string) {
	app.Get(PING_PATH, func(c *fiber.Ctx) error {
		return c.SendString(message)
	})
}

func setVersion(app *fiber.App, message string) {
	app.Get(VERSION_PATH, func(c *fiber.Ctx) error {
		return c.SendString(message)
	})
}

func Run(app *fiber.App, port string) {
	log.Fatalln(app.Listen(fmt.Sprintf(":%s", port)))
}
