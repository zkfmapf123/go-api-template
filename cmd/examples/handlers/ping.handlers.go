package handlers

import "github.com/gofiber/fiber/v2"

type PingResponse struct {
	Msg string
}

// PingHandlers godoc
//
//	@summary		Send ping
//	@description	ping
//	@accept			json
//	@param			Authorization	header		string	false	"인증 키"
//	@success		200				{object}	handlers.PingResponse
//	@router			/api/ping [get]
func PingHandlers(c *fiber.Ctx) error {
	return c.JSON(PingResponse{Msg: "hello world"})
}
