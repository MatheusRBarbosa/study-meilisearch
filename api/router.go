package api

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/api/v1/users"
	"squad10x.com.br/boilerplate/api/v1/utils"
)

func ApiRouter(server *fiber.App) {
	router := server.Group("/api")

	v1 := router.Group("/v1")
	utils.RegisterUtilsRoutes(v1)
	users.RegisterUsersRoutes(v1)
}
