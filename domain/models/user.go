package models

import (
	"time"

	"gorm.io/gorm"
	"squad10x.com.br/boilerplate/domain/dtos"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *User) ParseDto() dtos.UserDto {
	return dtos.UserDto{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
