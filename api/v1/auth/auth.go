package auth

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/api/requests"
	"squad10x.com.br/boilerplate/application/handlers"
)

// @Summary Login
// @Description Authentication
// @Tags User
// @Router /api/v1/login [post]
// @Param request body requests.LoginRequest true "Payload"
// @Success 200 {object} handlers.jwt
// @Failure 400 {object} exceptions.DomainError
func handleAuth(ctx *fiber.Ctx) error {
	req := new(requests.LoginRequest)

	err := requests.ValidateRequest(ctx, req)
	if err != nil {
		return nil
	}

	handler := handlers.AuthHandler()
	res, err := handler.Login(req.Email, req.Password)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(200).JSON(res)
	return nil
}
