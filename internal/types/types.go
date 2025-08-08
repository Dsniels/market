package types

import (
	"time"
)

type base struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"-"`
}

type Sucursal struct {
	base
	Nombre    string `json:"nombre" gorm:"not null"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Activo    bool   `json:"activo"`
}

type Categoria struct { 
	base
	Nombre string `json:"nombre" gorm:"not null"`
}

type Producto struct {
	base
	Nombre string `json:"nombre" gorm:"not null"`
	CategoriaID uint `gorm:"index"`
	Precio float32 `gorm:"type:numeric(8,2); not null"`
	Imagen string 
	Stock bool

}

type FormaPago struct { 
	base
	Nombre string `gorm:"not null; unique"`
}

type ProductoCompuesto struct {
	base
	ProductoPrincipalID uint
	ProductoComponenteID uint
}