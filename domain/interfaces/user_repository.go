package interfaces

import m "squad10x.com.br/boilerplate/domain/models"

type UserRepository interface {
	Create(u m.User) m.User
	GetByEmail(email string, j Join) (m.User, error)
	GetById(id int) (m.User, error)
	Save(u *m.User)
	SaveConfirmation(c *m.Confirmation)
}
