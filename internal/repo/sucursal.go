package repo

import "github.com/dsniels/market/internal/types"

type ISucursal interface{}
type Sucursal struct {
	*GenericRepo[types.Sucursal]
}
