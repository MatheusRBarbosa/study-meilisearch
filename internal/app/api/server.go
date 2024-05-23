package api

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	_ "squad10x.com.br/boilerplate/docs"
	"squad10x.com.br/boilerplate/internal/pkg/db"
	"squad10x.com.br/boilerplate/pkg/configs"
)

func StartHttpServer() {
	server := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	server.Get("/docs/*", swagger.HandlerDefault)

	server.Use(cors.New())

	db.ConnectDatabase()
	ApiRouter(server)

	var url string
	if configs.IsLocal() {
		url = "localhost:3000"
	} else {
		url = fmt.Sprintf(":%s", configs.GetEnv("APP_PORT"))
	}

	server.Listen(url)
}
