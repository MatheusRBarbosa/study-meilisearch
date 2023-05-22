package users

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/api/requests"
	"squad10x.com.br/boilerplate/application/handlers"
)

// @Summary User signup
// @Description Create a new user
// @Tags User
// @Router /api/v1/users [post]
// @Param request body requests.SignupRequest true "Payload"
// @Success 201 {object} dtos.UserDto
// @Failure 400 {object} exceptions.DomainError
func handleSignup(ctx *fiber.Ctx) error {
	req := new(requests.SignupRequest)

	err := requests.ValidateRequest(ctx, req)
	if err != nil {
		return nil
	}

	user := req.ParseToUser()
	handler := handlers.UserHandler()
	res, err := handler.Create(user)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(201).JSON(res)
	return nil
}
