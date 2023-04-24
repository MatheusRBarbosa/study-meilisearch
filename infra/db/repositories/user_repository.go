package repositories

import (
	"gorm.io/gorm"
	i "squad10x.com.br/boilerplate/domain/interfaces"
	m "squad10x.com.br/boilerplate/domain/models"
	"squad10x.com.br/boilerplate/infra/db"
)

type userRepository struct {
	ctx *gorm.DB
}

func UserRepository() i.UserRepository {
	return &userRepository{
		ctx: db.Context(),
	}
}

func (r *userRepository) Create(user m.User) m.User {
	result := r.ctx.Create(&user)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	return user
}

func (r *userRepository) GetByEmail(email string) (m.User, error) {
	user := m.User{}

	err := r.ctx.Where(&m.User{Email: email}).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, gorm.ErrRecordNotFound
	}

	return user, err
}

func (r *userRepository) GetById(id int) (m.User, error) {
	user := m.User{}
	err := r.ctx.Where(&m.User{ID: id}).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, gorm.ErrRecordNotFound
	}

	return user, err
}
