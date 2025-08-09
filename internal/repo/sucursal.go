package repo

import (
	"github.com/dsniels/market/core/types"
	"gorm.io/gorm"
)

type ISucursal interface{}
type Sucursal struct {
	*Generic[types.Sucursal]
}

func NewSucursal(db *gorm.DB) *Sucursal {
	g := NewGeneric[types.Sucursal](db)
	return &Sucursal{
		Generic: g,
	}
}
