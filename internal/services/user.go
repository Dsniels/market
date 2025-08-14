package services

import (
	"context"
	"fmt"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/repo"
)

type IUser interface {
	Create(context.Context, *types.User) error
	GetAll(context.Context) (*[]types.User, error)
	GetById(context.Context, uint) (*types.User, error)
	GetByEmail(context.Context, string) (*types.User, error)
	Update(context.Context, uint, *types.User) (*types.User, error)
	UpdatePassword(context.Context, string, string, uint) error
	Delete(context.Context, uint) error
}

type User struct {
	repo repo.IUser
}

func (u *User) Create(ctx context.Context, user *types.User) error {
	return u.repo.Create(ctx, &types.User{})
}

func (u *User) GetAll(ctx context.Context) (*[]types.User, error) {
	users, err := u.repo.GetList(ctx)
	if err != nil {
		return nil, nil
	}
	return users, err
}

func (u *User) GetById(ctx context.Context, id uint) (*types.User, error) {
	user, err := u.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) GetByEmail(ctx context.Context, email string) (*types.User, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) UpdatePassword(ctx context.Context, currentPwd string, newPwd string, id uint) error {
	user, err := u.GetById(ctx, id)
	if err != nil {
		return err
	}
	if ok, _ := user.ValidatePassword(currentPwd); !ok {
		return fmt.Errorf("incorrect password")
	}

	err = user.SetPassword(newPwd)
	if err != nil {
		return err
	}
	_, err = u.repo.Update(user)
	return err
}

func (u *User) Update(ctx context.Context, id uint, user *types.User) (*types.User, error) {
	_, err := u.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	user.ID = id
	newUser, err := u.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *User) Delete(ctx context.Context, id uint) error {
	return u.repo.Delete(ctx, id)
}

func NewUser(repo repo.IUser) *User {
	return &User{
		repo: repo,
	}
}
