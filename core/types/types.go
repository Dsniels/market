package types

import (
	"time"
)

type Base struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"-"`
}

type Categoria struct {
	Base   `gorm:"embedded"`
	Nombre string `json:"nombre" gorm:"not null; unique"`
}
type Sucursal struct {
	Base      `gorm:"embedded"`
	Nombre    string `json:"nombre" gorm:"not null"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Activo    bool   `json:"activo"`
}

type Producto struct {
	Base        `gorm:"embedded"`
	Nombre      string  `json:"nombre" gorm:"not null"`
	CategoriaID uint    `gorm:"index"`
	Precio      float32 `gorm:"type:numeric(8,2); not null"`
	Imagen      string
	Stock       bool
}

type FormaPago struct {
	Base   `gorm:"embedded"`
	Nombre string `gorm:"not null; unique"`
}

type ProductoCompuesto struct {
	Base                 `gorm:"embedded"`
	ProductoPrincipalID  uint
	ProductoComponenteID uint
}
