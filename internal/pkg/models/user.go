package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"squad10x.com.br/boilerplate/internal/pkg/enums"
)

type UserCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	RoleId    int `gorm:"column:roleId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	// Relations
	Role         Role
	Confirmation Confirmation
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now().UTC()
	u.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now().UTC()
	return nil
}

func (u *User) IsSadmin() bool {
	return u.RoleId == int(enums.Sadmin)
}

func (u *User) IsAdmin() bool {
	return u.RoleId == int(enums.Sadmin) ||
		u.RoleId == int(enums.Admin)
}

func (u *User) IsUser() bool {
	return u.RoleId == int(enums.User)
}
