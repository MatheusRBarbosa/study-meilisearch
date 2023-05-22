package api

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"squad10x.com.br/boilerplate/api/middlewares"
	"squad10x.com.br/boilerplate/crosscutting"
	"squad10x.com.br/boilerplate/infra/db"

	_ "squad10x.com.br/boilerplate/docs"
)

func StartHttpServer() {
	server := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: middlewares.ErrorHandler(),
	})

	server.Get("/docs/*", swagger.HandlerDefault)

	server.Use(cors.New())

	db.ConnectDatabase()
	ApiRouter(server)

	var url string
	if crosscutting.IsLocal() {
		url = "localhost:3000"
	} else {
		url = ":3000"
	}

	server.Listen(url)
}
