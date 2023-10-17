package utils

import "github.com/gofiber/fiber/v2"

// @Summary API health check
// @Tags Utils
// @Router /api/v1/health [get]
// @Success 200 {object} object
func handleHealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "healthy",
	})
}
