package interfaces

import (
	"github.com/golang-jwt/jwt/v5"
	m "squad10x.com.br/boilerplate/domain/models"
)

type AuthService interface {
	Generate(user m.User) string
	Validate(token string) (*jwt.Token, error)
	SetAuthUser(user m.User) error
	GetAuthUser() *m.User
}
