package handlers

import (
	"github.com/gofiber/fiber/v2"
	"squad10x.com.br/boilerplate/crosscutting/logger"
	"squad10x.com.br/boilerplate/domain/dtos"
	"squad10x.com.br/boilerplate/domain/enums"
	i "squad10x.com.br/boilerplate/domain/interfaces"
	m "squad10x.com.br/boilerplate/domain/models"
	"squad10x.com.br/boilerplate/domain/services"
	"squad10x.com.br/boilerplate/infra/db/repositories"
)

type userHandler struct {
	logger         logger.Logger
	userRepository i.UserRepository
	userService    i.UserService
}

func UserHandler() *userHandler {
	return &userHandler{
		logger:         logger.GetLogger(),
		userRepository: repositories.UserRepository(),
		userService:    services.UserService(),
	}
}

func (h *userHandler) HandleCreate(u m.User) (dtos.UserDto, error) {
	u.RoleId = int(enums.User)
	u = h.userRepository.Create(u)
	return u.ParseDto(), nil
}

func (h *userHandler) HandleForgotPassword(email string) (dtos.UserDto, error) {
	user, _ := h.userRepository.GetByEmail(email, i.Preload{
		Rels: []string{"Confirmation"},
	})

	user.Confirmation.Confirmed = false
	user.Confirmation.Code = h.userService.GenerateCode()
	user.Confirmation.UserId = user.ID

	h.userRepository.UpsertConfirmation(user.Confirmation)
	// TODO: Send email confirmation

	return user.ParseDto(), nil
}

func (h *userHandler) HandleValidateCode(email, code string) (fiber.Map, error) {
	user, _ := h.userRepository.GetByEmail(email, i.Preload{Rels: []string{"Confirmation"}})

	res, err := h.userService.ValidateCode(code, user)
	if err != nil {
		return fiber.Map{}, err
	}

	return fiber.Map{"status": res}, nil
}
