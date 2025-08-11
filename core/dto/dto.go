package dto

type ProductoComp struct {
	ID                  uint `json:"id"`
	ProductoPrincipalID uint
	ProductoComponenteID uint
}
