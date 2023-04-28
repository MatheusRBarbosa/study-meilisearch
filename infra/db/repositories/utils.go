package repositories

import (
	"gorm.io/gorm"
	i "squad10x.com.br/boilerplate/domain/interfaces"
)

func PreloadRelations(tx *gorm.DB, p i.Preload) *gorm.DB {
	for _, r := range p.Rels {
		tx.Preload(r)
	}

	return tx
}
