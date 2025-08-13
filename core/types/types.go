package types

import (
	"time"

	"golang.org/x/crypto/bcrypt"
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

type Terminal struct {
	Base
	Nombre     string
	SucursalID uint
	Sucursal   Sucursal `gorm:"foreignKey:SucursalID;referces:ID"`
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
	ProductoPrincipal    Producto `gorm:"foreignKey:ProductoPrincipalID;references:ID"`
	ProductoComponenteID uint
	ProductoComponente   Producto `gorm:"foreignKey:ProductoComponenteID;references:ID"`
}

type User struct {
	Base
	Nombre    string `gorm:"not null"`
	Apellido  string `gorm:"not null"`
	Telefono  string
	Direccion string
	Email     string `gorm:"not null; unique;"`
	Hash      []byte `gorm:"not null" json:"-"`
}

func (u *User) SetPassword(plain string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), 16)
	if err != nil {
		return err
	}
	u.Hash = hash
	return nil
}

func (u *User) ValidatePassword(plain string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(u.Hash, []byte(plain))
	if err != nil {
		return false, err
	}

	return true, nil
}
