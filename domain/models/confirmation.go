package models

import (
	"time"

	"gorm.io/gorm"
)

type Confirmation struct {
	ID        int
	UserId    int `gorm:"column:userId"`
	Code      string
	Confirmed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Confirmation) BeforeCreate(tx *gorm.DB) error {
	c.CreatedAt = time.Now().UTC()
	c.UpdatedAt = time.Now().UTC()
	return nil
}

func (c *Confirmation) BeforeUpdate(tx *gorm.DB) error {
	c.UpdatedAt = time.Now()
	return nil
}
