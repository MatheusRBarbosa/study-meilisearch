package repositories

import (
	"gorm.io/gorm"
	"squad10x.com.br/boilerplate/internal/pkg/utils"
)

func LoadRelations(tx *gorm.DB, p utils.Join) *gorm.DB {
	for _, r := range p.Rels {
		tx.Joins(r)
	}

	return tx
}
