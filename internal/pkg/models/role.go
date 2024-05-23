package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	r.CreatedAt = time.Now().UTC()
	r.UpdatedAt = time.Now().UTC()
	return nil
}

func (r *Role) BeforeUpdate(tx *gorm.DB) error {
	r.UpdatedAt = time.Now().UTC()
	return nil
}
