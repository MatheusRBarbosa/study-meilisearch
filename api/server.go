package api

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/infra/db"
)

func StartHttpServer() {
	server := fiber.New()

	db.ConnectDatabase()
	ApiRouter(server)

	server.Listen(":3000")
}
