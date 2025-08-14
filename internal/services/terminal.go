package services

import (
	"context"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/repo"
)

type ITerminal interface {
	GetAll(ctx context.Context) (*[]types.Terminal, error)
	GetByID(ctx context.Context, id uint) (*types.Terminal, error)
	Create(ctx context.Context, terminal *types.Terminal) error
	Update(ctx context.Context, id uint, terminal *types.Terminal) error
	DeleteByID(ctx context.Context, id uint) error
}

type Terminal struct {
	repo repo.ITerminal
}

func (t *Terminal) GetAll(ctx context.Context) (*[]types.Terminal, error) {
	terminals, err := t.repo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return terminals, nil

}

func (t *Terminal) GetByID(ctx context.Context, id uint) (*types.Terminal, error) {
	terminal, err := t.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return terminal, nil
}

func (t *Terminal) Create(ctx context.Context, terminal *types.Terminal) error {
	err := t.repo.Create(ctx, terminal)
	if err != nil {
		return err
	}
	return err
}

func (t *Terminal) Update(ctx context.Context, id uint, terminal *types.Terminal) error {
	cur, err := t.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	terminal.ID = cur.ID
	_, err = t.repo.Update(terminal)
	return err

}

func (t *Terminal) DeleteByID(ctx context.Context, id uint) error {
	_, err := t.GetByID(ctx, id)
	if err != nil {
		return err
	}
	err = t.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func NewTerminal(repo repo.ITerminal) *Terminal {
	return &Terminal{
		repo: repo,
	}
}
