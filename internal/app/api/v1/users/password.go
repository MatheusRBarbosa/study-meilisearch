package users

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/internal/app/api/requests"
	"squad10x.com.br/boilerplate/internal/pkg/handlers"
)

// @Summary Forgot user password
// @Description First step to user recovery password
// @Tags User
// @Router /api/v1/forgot-password [post]
// @Param request body requests.ForgotPasswordRequest true "Payload"
// @Success 200 {object} dtos.UserDto
// @Failure 400 {object} exceptions.DomainError
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

// @Summary Validate code
// @Description Second step to user recovery password
// @Tags User
// @Router /api/v1/validate-code [post]
// @Param request body requests.ValidateCodeRequest true "Payload"
// @Success 200 {object} object
// @Failure 400 {object} exceptions.DomainError
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

// @Summary Change user password
// @Description Final step to user recovery password
// @Tags User
// @Router /api/v1/change-password [post]
// @Param request body requests.ChangePasswordRequest true "Payload"
// @Success 200 {object} dtos.UserDto
// @Failure 400 {object} exceptions.DomainError
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
