package handlers

import (
	"golang.org/x/crypto/bcrypt"
	"squad10x.com.br/boilerplate/domain/exceptions"
	i "squad10x.com.br/boilerplate/domain/interfaces"
	"squad10x.com.br/boilerplate/domain/services"
	"squad10x.com.br/boilerplate/infra/db/repositories"
)

type jwt struct {
	Token string `json:"token"`
}

type authHandler struct {
	userRepository i.UserRepository
	authService    i.AuthService
}

func AuthHandler() *authHandler {
	return &authHandler{
		userRepository: repositories.UserRepository(),
		authService:    services.AuthService(),
	}
}

func (h *authHandler) HandleLogin(email, password string) (jwt, error) {
	res := jwt{}
	user, err := h.userRepository.GetByEmail(email)
	invalidPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil || invalidPassword != nil {
		return res, exceptions.UNAUTHORIZED
	}

	return jwt{
		Token: h.authService.Generate(user),
	}, nil
}
