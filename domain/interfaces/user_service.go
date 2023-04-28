package interfaces

import m "squad10x.com.br/boilerplate/domain/models"

type UserService interface {
	GenerateCode() string
	ValidateCode(code string, u m.User) (bool, error)
}
