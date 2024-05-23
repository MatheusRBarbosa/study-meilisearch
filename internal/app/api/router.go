package api

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/internal/app/api/v1/pokemons"
)

func ApiRouter(server *fiber.App) {
	router := server.Group("/api")

	v1 := router.Group("/v1")
	pokemons.RegisterUtilsRoutes(v1)

}
