package handlers

import (
	"gorm.io/gorm"
	"squad10x.com.br/boilerplate/application/requests"
	"squad10x.com.br/boilerplate/crosscutting/logger"
	"squad10x.com.br/boilerplate/domain/dtos"
	"squad10x.com.br/boilerplate/domain/exceptions"
	i "squad10x.com.br/boilerplate/domain/interfaces"
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

func (h *userHandler) HandleCreate(req *requests.SignupRequest) (dtos.UserDto, error) {
	user := req.ParseToUser()

	_, err := h.userRepository.GetByEmail(user.Email)
	if err != nil && err == gorm.ErrRecordNotFound {
		user = h.userRepository.Create(user)
		return user.ParseDto(), nil
	} else {
		return user.ParseDto(), exceptions.EMAIL_ALREADY_EXIST
	}
}
