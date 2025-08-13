package repo

import (
	"context"

	"github.com/dsniels/market/core/types"
	"gorm.io/gorm"
)

type IProductoComp interface {
	IGeneric[types.ProductoCompuesto]
}

type ProductoComp struct {
	*Generic[types.ProductoCompuesto]
}

func (p *ProductoComp) GetById(ctx context.Context, id uint) (*types.ProductoCompuesto, error) {
	prodComp := new(types.ProductoCompuesto)
	err := p.db.Where("id = ?", id).Preload("ProductoPrincipal").Preload("ProductoComponente").First(prodComp).Error
	if err != nil {
		return nil, err
	}

	return prodComp, nil
}
func (p *ProductoComp) GetList(ctx context.Context) (*[]types.ProductoCompuesto, error) {
	prodComp := new([]types.ProductoCompuesto)
	err := p.db.Preload("ProductoPrincipal").Preload("ProductoComponente").Find(prodComp).Error
	if err != nil {
		return nil, err
	}
	return prodComp, nil
}

func NewProductoComp(db *gorm.DB) *ProductoComp {

	g := NewGeneric[types.ProductoCompuesto](db)
	return &ProductoComp{
		Generic: g,
	}

}
