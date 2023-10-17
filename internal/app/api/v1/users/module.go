package users

import "github.com/gofiber/fiber/v2"

func RegisterUsersRoutes(r fiber.Router) {
	r.Post("/users", handleSignup)

	r.Post("/forgot-password", handleForgotPassword)
	r.Post("/validate-code", handleValidateCode)
	r.Post("/change-password", handleChangePassword)
}
