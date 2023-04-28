package users

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/api/requests"
	"squad10x.com.br/boilerplate/application/handlers"
)

func handleForgotPassword(ctx *fiber.Ctx) error {
	req := new(requests.ForgotPasswordRequest)

	err := requests.ValidateRequest(ctx, req)
	if err != nil {
		return nil
	}

	handler := handlers.UserHandler()
	res, err := handler.ForgotPassword(req.Email)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(200).JSON(res)
	return nil
}

func handleValidateCode(ctx *fiber.Ctx) error {
	req := new(requests.ValidateCodeRequest)

	err := requests.ValidateRequest(ctx, req)
	if err != nil {
		return nil
	}

	handler := handlers.UserHandler()
	res, err := handler.ValidateCode(req.Email, req.Code)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(200).JSON(res)
	return nil
}

func handleChangePassword(ctx *fiber.Ctx) error {
	req := new(requests.ChangePasswordRequest)

	err := requests.ValidateRequest(ctx, req)
	if err != nil {
		return nil
	}

	handler := handlers.UserHandler()
	res, err := handler.ChangePassword(req.Code, req.Password, req.Email)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(200).JSON(res)
	return nil
}
