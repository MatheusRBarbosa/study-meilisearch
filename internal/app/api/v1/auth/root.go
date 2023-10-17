package auth

import "github.com/gofiber/fiber/v2"

func RegisterAuthRoutes(r fiber.Router) {
	r.Post("login", handleAuth)
}
