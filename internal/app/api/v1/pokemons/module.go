package pokemons

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterUtilsRoutes(r fiber.Router) {
	r.Get("/health", handleHealthCheck)
	r.Get("/search", handleSearch)
}
