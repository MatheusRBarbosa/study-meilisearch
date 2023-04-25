package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/domain/exceptions"
	m "squad10x.com.br/boilerplate/domain/models"
	"squad10x.com.br/boilerplate/domain/services"
	"squad10x.com.br/boilerplate/infra/db/repositories"
)

func ValidateJWT() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		const SCHEMA = "Bearer"
		header := ctx.Get("Authorization")
		if header == "" {
			ctx.Status(exceptions.UNAUTHORIZED.Code)
			ctx.JSON(exceptions.UNAUTHORIZED)
			return nil
		}

		jwt := strings.Trim(header[len(SCHEMA):], " ")
		authService := services.AuthService()
		token, _ := authService.Validate(jwt)
		if token.Valid {
			claims := token.Claims.(*m.UserCustomClaims)
			userRepository := repositories.UserRepository()

			user, err := userRepository.GetById(claims.ID)
			if err != nil {
				ctx.Status(exceptions.UNAUTHORIZED.Code)
				ctx.JSON(exceptions.UNAUTHORIZED)
			}

			authService.SetAuthUser(user)
		} else {
			ctx.Status(exceptions.UNAUTHORIZED.Code)
			ctx.JSON(exceptions.UNAUTHORIZED)
		}

		return ctx.Next()
	}
}
