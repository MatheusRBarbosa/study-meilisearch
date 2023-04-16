package utils

import "github.com/gofiber/fiber/v2"

func RegisterUtilsRoutes(g fiber.Router) {
	g.Get("/health", handleHealthCheck)
}
