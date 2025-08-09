package repo

import (
	"github.com/dsniels/market/core/types"
	"gorm.io/gorm"
)

type ICategoria interface {
	IGeneric[types.Categoria]
}

type Categoria struct {
	db *gorm.DB
	*Generic[types.Categoria]
}

func NewCategoria(db *gorm.DB) *Categoria {
	g := NewGeneric[types.Categoria](db)
	return &Categoria{
		db:      db,
		Generic: g,
	}
}
