package users

import "github.com/gofiber/fiber/v2"

func RegisterUsersRoutes(r fiber.Router) {
	r.Post("users", handleSignup)
}
