package repo

import (
	"github.com/dsniels/market/core/types"
	"gorm.io/gorm"
)

type ITerminal interface {
	IGeneric[types.Terminal]
}

type Terminal struct {
	*Generic[types.Terminal]
}

func NewTerminal(db *gorm.DB) *Terminal {
	g := NewGeneric[types.Terminal](db)
	return &Terminal{
		Generic: g,
	}
}
