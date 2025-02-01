package apis

import (
	"github.com/gofiber/fiber/v2"
)

func HealthCheck() fiber.Handler {

	return func (c *fiber.Ctx) error  {
		return c.SendString("Hello world")
	}
}