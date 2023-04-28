package requests

import (
	"golang.org/x/crypto/bcrypt"
	l "squad10x.com.br/boilerplate/crosscutting/logger"
	m "squad10x.com.br/boilerplate/domain/models"
)

type SignupRequest struct {
	Name     string `json:"name" validate:"required,lte=255"`
	Email    string `json:"email" validate:"required,email,lte=255,unique=users"`
	Password string `json:"password" validate:"required,gte=6,lte=255"`
}

func (r *SignupRequest) ParseToUser() m.User {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		l.GetLogger().Panicln(err)
	}

	return m.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: string(encryptedPassword),
	}
}
