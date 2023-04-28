package repositories

import (
	"gorm.io/gorm"
	i "squad10x.com.br/boilerplate/domain/interfaces"
)

func LoadRelations(tx *gorm.DB, p i.Join) *gorm.DB {
	for _, r := range p.Rels {
		tx.Joins(r)
	}

	return tx
}
