package repo

import "github.com/dsniels/market/internal/types"

type IProductoComp interface{}

type ProductoComp struct {
	*GenericRepo[types.ProductoCompuesto]
}
