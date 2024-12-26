package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware() fiber.Handler {

	return func(c *fiber.Ctx) error {
		log.Printf("Method : %s Path : %s", c.Method(),c.Path() )
		err := c.Next()
		return err
	}
}