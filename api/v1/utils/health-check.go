package utils

import "github.com/gofiber/fiber/v2"

func handleHealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "healthy",
	})
}
