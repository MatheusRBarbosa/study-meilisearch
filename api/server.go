package api

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/crosscutting"
	"squad10x.com.br/boilerplate/infra/db"
)

func StartHttpServer() {
	server := fiber.New()

	db.ConnectDatabase()
	ApiRouter(server)

	url := crosscutting.GetEnv("APP_URL")
	server.Listen(url)
}
