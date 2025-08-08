package repo

import (
	"github.com/dsniels/market/internal/types"
	"gorm.io/gorm"
)

type ICategoriaRepo interface {
	IGenericRepo[types.Categoria]
}

type CategoriaRepo struct {
	db *gorm.DB
	GenericRepo[types.Categoria]
}

func NewCategoriaRepo(db *gorm.DB) *CategoriaRepo {
	return &CategoriaRepo{
		db: db,
	}
}
