package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	ex "squad10x.com.br/boilerplate/domain/exceptions"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := ex.DEFAULT.Code

		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		err = ctx.Status(code).JSON(err)
		if err != nil {
			return ctx.Status(ex.DEFAULT.Code).JSON(ex.DEFAULT)
		}

		return nil
	}
}
