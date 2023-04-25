package api

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"squad10x.com.br/boilerplate/api/middlewares"
	"squad10x.com.br/boilerplate/crosscutting"
	"squad10x.com.br/boilerplate/infra/db"
)

func StartHttpServer() {
	server := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: middlewares.ErrorHandler(),
	})

	server.Use(cors.New())

	db.ConnectDatabase()
	ApiRouter(server)

	url := crosscutting.GetEnv("APP_URL")
	server.Listen(url)
}
