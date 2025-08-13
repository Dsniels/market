package repo

import "github.com/dsniels/market/core/types"

type IUser interface {
	IGeneric[types.User]
}

type User struct {
	*Generic[types.User]
}

func NewUser() *User {
	return &User{}
}
