package dto

import "github.com/dsniels/market/core/types"

type ProductoComp struct {
	ID                uint           `json:"id"`
	ProductoPrincipal *types.Producto `json:"producto_principal"`
	ProductoCompuesto *types.Producto `json:"producto_compuesto"`
}
