package auth

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/api/requests"
	"squad10x.com.br/boilerplate/application/handlers"
)

func handleAuth(ctx *fiber.Ctx) error {
	req := new(requests.LoginRequest)

	err := requests.ValidateReqeust(ctx, req)
	if err != nil {
		return nil
	}

	handler := handlers.AuthHandler()
	res, err := handler.HandleLogin(req.Email, req.Password)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(200)
	ctx.JSON(res)
	return nil
}
