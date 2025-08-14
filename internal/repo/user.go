package repo

import (
	"context"

	"github.com/dsniels/market/core/types"
)

type IUser interface {
	GetByEmail(context.Context, string) (*types.User, error)
	IGeneric[types.User]
}

type User struct {
	*Generic[types.User]
}

func (u *User) GetByEmail(ctx context.Context, email string) (*types.User, error) {
	user := new(types.User)
	err := u.db.Where("email = ?", email).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUser() *User {
	return &User{}
}
