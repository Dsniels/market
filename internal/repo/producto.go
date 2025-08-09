package repo

import (
	"github.com/dsniels/market/core/types"
	"gorm.io/gorm"
)

type IProducto interface {
	IGeneric[types.Producto]
}
type Producto struct {
	*Generic[types.Producto]
}

func NewProducto(db *gorm.DB) *Producto {

	g := NewGeneric[types.Producto](db)
	return &Producto{
		Generic: g,
	}
}
