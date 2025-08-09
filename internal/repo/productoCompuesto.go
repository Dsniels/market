package repo

import (
	"github.com/dsniels/market/core/types"
	"gorm.io/gorm"
)

type IProductoComp interface {
	IGeneric[types.ProductoCompuesto]
}

type ProductoComp struct {
	*Generic[types.ProductoCompuesto]
}

func NewProductoComp(db *gorm.DB) *ProductoComp {

	g := NewGeneric[types.ProductoCompuesto](db)
	return &ProductoComp{
		Generic: g,
	}

}
