package api

import "github.com/gofiber/fiber/v2"

func StartHttpServer() {
	server := fiber.New()

	ApiRouter(server)

	server.Listen(":3000")
}
