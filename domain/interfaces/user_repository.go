package interfaces

import m "squad10x.com.br/boilerplate/domain/models"

type UserRepository interface {
	Create(u m.User) m.User
	GetByEmail(email string) (m.User, error)
	GetById(id int) (m.User, error)
}