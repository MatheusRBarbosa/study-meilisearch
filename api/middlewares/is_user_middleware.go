package middlewares

import (
	"github.com/gofiber/fiber/v2"
	ex "squad10x.com.br/boilerplate/domain/exceptions"
	"squad10x.com.br/boilerplate/domain/services"
)

func IsUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authService := services.AuthService()

		authorized := authService.AuthedIsUser()
		if !authorized {
			ctx.Status(ex.FORBIDDEN.Code).JSON(ex.FORBIDDEN)
			return nil
		}

		return ctx.Next()
	}
}
