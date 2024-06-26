package pokemons

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/internal/pkg/handlers"
)

// @Summary API health check
// @Tags Utils
// @Router /api/v1/health [get]
// @Success 200 {object} object
func handleHealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "healthy",
	})
}

func handleSearch(ctx *fiber.Ctx) error {
	name := ctx.Query("name")

	handler := handlers.NewGeneralHandler()
	res, err := handler.Search(name)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(200).JSON(res)
	return nil
}
