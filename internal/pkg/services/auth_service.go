package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	m "squad10x.com.br/boilerplate/internal/pkg/models"
	"squad10x.com.br/boilerplate/pkg/configs"
	"squad10x.com.br/boilerplate/pkg/logger"
)

var authedUser *m.User

type AuthService interface {
	Generate(u m.User) string
	Validate(token string) (*jwt.Token, error)
	SetAuthUser(u m.User)
	GetAuthUser() *m.User
	AuthedIsSadmin() bool
	AuthedIsAdmin() bool
	AuthedIsUser() bool
}

type authService struct {
	secret string
	issure string
	logger logger.Logger
}

func NewAuthService() AuthService {
	return &authService{
		secret: configs.GetEnv("JWT_SECRET"),
		issure: "squad10x.com.br",
		logger: logger.GetLogger(),
	}
}

func (s *authService) GetAuthUser() *m.User {
	return authedUser
}

func (s *authService) SetAuthUser(u m.User) {
	authedUser = &u
}

func (s *authService) Generate(user m.User) string {
	claims := &m.UserCustomClaims{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7 * 2)),
			Issuer:    s.issure,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		s.logger.Panicln(err)
	}

	return jwtToken
}

func (s *authService) Validate(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &m.UserCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])

		}
		return []byte(s.secret), nil
	})
}

func (s *authService) AuthedIsSadmin() bool {
	if authedUser == nil {
		return false
	}

	return authedUser.IsSadmin()
}

func (s *authService) AuthedIsAdmin() bool {
	if authedUser == nil {
		return false
	}

	return authedUser.IsAdmin()
}

func (s *authService) AuthedIsUser() bool {
	if authedUser == nil {
		return false
	}

	return authedUser.IsUser()
}
