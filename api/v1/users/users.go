package users

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/application/handlers"
	"squad10x.com.br/boilerplate/application/requests"
	"squad10x.com.br/boilerplate/crosscutting/logger"
)

func handleSignup(ctx *fiber.Ctx) error {
	req := new(requests.SignupRequest)
	l := logger.GetLogger()

	if err := ctx.BodyParser(req); err != nil {
		l.Error(err)
	}

	if err := requests.GetValidator().Struct(req); err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		ctx.JSON(fiber.Map{"errors": requests.ParseError(err)})
		return nil
	}

	handler := handlers.UserHandler()
	res, err := handler.HandleCreate(req)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(201)
	ctx.JSON(res)
	return nil
}
