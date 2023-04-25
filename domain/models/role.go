package models

import (
	"time"

	"gorm.io/gorm"
	"squad10x.com.br/boilerplate/domain/dtos"
)

type Role struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (r *Role) ParseDto() dtos.IdNameDto {
	return dtos.IdNameDto{
		Id:   r.ID,
		Name: r.Name,
	}
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	return nil
}

func (r *Role) BeforeUpdate(tx *gorm.DB) error {
	r.UpdatedAt = time.Now()
	return nil
}