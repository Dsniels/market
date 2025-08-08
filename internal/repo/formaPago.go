package repo

import "github.com/dsniels/market/internal/types"

type IFormaPago interface{}

type FormaPago struct {
	*GenericRepo[types.FormaPago]
}
