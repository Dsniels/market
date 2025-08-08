package repo

import "github.com/dsniels/market/internal/types"

type IProducto interface{}
type Producto struct {
	*GenericRepo[types.Producto]
}
