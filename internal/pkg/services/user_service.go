package services

import (
	"math/rand"
	"strconv"
	"time"

	ex "squad10x.com.br/boilerplate/internal/pkg/exceptions"
	m "squad10x.com.br/boilerplate/internal/pkg/models"
)

type userService struct{}

type UserService interface {
	GenerateCode() string
	ValidateCode(code string, u m.User) (bool, error)
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) GenerateCode() string {
	code := strconv.Itoa(rand.Intn(9999))
	for len(code) < 5 {
		code = "0" + code
	}

	return code
}

func (s *userService) ValidateCode(code string, u m.User) (bool, error) {
	if code != u.Confirmation.Code {
		return false, ex.CODE_INVALID
	}

	if time.Now().Sub(u.Confirmation.UpdatedAt).Minutes() > 15 {
		return false, ex.CODE_EXPIRED
	}

	if u.Confirmation.Confirmed {
		return false, ex.CODE_ALREADY_USED
	}

	return true, nil
}
