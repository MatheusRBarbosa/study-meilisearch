package handlers

import (
	"golang.org/x/crypto/bcrypt"
	"squad10x.com.br/boilerplate/internal/pkg/db/repositories"
	"squad10x.com.br/boilerplate/internal/pkg/exceptions"
	"squad10x.com.br/boilerplate/internal/pkg/services"
	"squad10x.com.br/boilerplate/internal/pkg/utils"
)

type jwt struct {
	Token string `json:"token"`
}

type authHandler struct {
	userRepository repositories.UserRepository
	authService    services.AuthService
}

func AuthHandler() *authHandler {
	return &authHandler{
		userRepository: repositories.NewUserRepository(),
		authService:    services.NewAuthService(),
	}
}

func (h *authHandler) Login(email, password string) (jwt, error) {
	res := jwt{}
	user, err := h.userRepository.GetByEmail(email, utils.Join{})
	invalidPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil || invalidPassword != nil {
		return res, exceptions.UNAUTHORIZED
	}

	return jwt{
		Token: h.authService.Generate(user),
	}, nil
}
