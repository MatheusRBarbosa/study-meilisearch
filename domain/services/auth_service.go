package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"squad10x.com.br/boilerplate/crosscutting"
	"squad10x.com.br/boilerplate/crosscutting/logger"
	i "squad10x.com.br/boilerplate/domain/interfaces"
	m "squad10x.com.br/boilerplate/domain/models"
)

var authedUser *m.User

type authService struct {
	secret string
	issure string
	logger logger.Logger
}

func AuthService() i.AuthService {
	return &authService{
		secret: crosscutting.GetEnv("JWT_SECRET"),
		issure: "squad10x.com.br",
		logger: logger.GetLogger(),
	}
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

func (s *authService) GetAuthUser() *m.User {
	return authedUser
}

func (s *authService) SetAuthUser(u m.User) error {
	authedUser = &u
	return nil
}
