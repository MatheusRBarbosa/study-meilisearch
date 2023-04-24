package handlers

import (
	"gorm.io/gorm"
	"squad10x.com.br/boilerplate/crosscutting/logger"
	"squad10x.com.br/boilerplate/domain/dtos"
	"squad10x.com.br/boilerplate/domain/exceptions"
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
	_, err := h.userRepository.GetByEmail(u.Email)

	if err != nil && err == gorm.ErrRecordNotFound {
		u = h.userRepository.Create(u)
		return u.ParseDto(), nil
	} else {
		return u.ParseDto(), exceptions.EMAIL_ALREADY_EXIST
	}
}
