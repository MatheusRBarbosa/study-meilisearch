package middlewares

import (
	"github.com/gofiber/fiber/v2"
	ex "squad10x.com.br/boilerplate/internal/pkg/exceptions"
	"squad10x.com.br/boilerplate/internal/pkg/services"
)

func IsAdmin() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authService := services.NewAuthService()

		authorized := authService.AuthedIsAdmin()
		if !authorized {
			ctx.Status(ex.FORBIDDEN.Code).JSON(ex.FORBIDDEN)
			return nil
		}

		return ctx.Next()
	}
}
