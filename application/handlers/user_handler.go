package handlers

import (
	"math/rand"
	"strconv"

	"squad10x.com.br/boilerplate/crosscutting/logger"
	"squad10x.com.br/boilerplate/domain/dtos"
	"squad10x.com.br/boilerplate/domain/enums"
	i "squad10x.com.br/boilerplate/domain/interfaces"
	m "squad10x.com.br/boilerplate/domain/models"
	"squad10x.com.br/boilerplate/infra/db/repositories"
)

type userHandler struct {
	logger         logger.Logger
	userRepository i.UserRepository
}

func UserHandler() *userHandler {
	return &userHandler{
		logger:         logger.GetLogger(),
		userRepository: repositories.UserRepository(),
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

	user.Confirmation.Confirmed = user.Confirmation.ID != 0
	user.Confirmation.Code = generateCode()
	user.Confirmation.UserId = user.ID

	h.userRepository.UpsertConfirmation(user.Confirmation)
	// TODO: Send email confirmation

	return user.ParseDto(), nil
}

func generateCode() string {
	code := strconv.Itoa(rand.Intn(9999))
	for len(code) < 5 {
		code = "0" + code
	}

	return code
}
