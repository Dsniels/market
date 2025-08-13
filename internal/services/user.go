package services

import (
	"context"

	"github.com/dsniels/market/core/types"
)

type IUser interface {
	Create(context.Context, *types.User) error
	GetAll(context.Context) (*[]types.User, error)
	GetById(context.Context, uint) (*types.User, error)
	Update(context.Context, uint, *types.User) (*types.User, error)
	Delete(context.Context, uint) (*types.User, error)
}

type User struct {
}

func (u *User) Create(context.Context, *types.User) error {
	return nil
}
func (u *User) GetAll(context.Context) (*[]types.User, error) {
	return nil, nil
}
func (u *User) GetById(context.Context, uint) (*types.User, error) {
	return nil, nil
}
func (u *User) Update(context.Context, uint, *types.User) (*types.User, error) {
	return nil, nil
}
func (u *User) Delete(context.Context, uint) (*types.User, error) {
	return nil, nil
}

func NewUser() *User {
	return &User{}
}
