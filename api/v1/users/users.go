package users

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/api/requests"
	"squad10x.com.br/boilerplate/crosscutting/logger"
)

func handleSignup(ctx *fiber.Ctx) error {
	request := new(requests.SignupRequest)
	l := logger.GetLogger()

	if err := ctx.BodyParser(request); err != nil {
		l.Error(err)
	}

	if err := requests.GetValidator().Struct(request); err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		ctx.JSON(fiber.Map{"errors": requests.ParseError(err)})
		return nil
	}

	ctx.Status(201)
	ctx.JSON(fiber.Map{"ok": "true"})
	return nil
}
