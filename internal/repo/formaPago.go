package repo

import (
	"github.com/dsniels/market/core/types"
	"gorm.io/gorm"
)

type IFormaPago interface{
	IGeneric[types.FormaPago]
}

type FormaPago struct {
	*Generic[types.FormaPago]
}

func NewFormaPago(db *gorm.DB) *FormaPago {
	g := NewGeneric[types.FormaPago](db)
	return &FormaPago{
		Generic: g,
	}

}
