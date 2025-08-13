package services

import (
	"context"

	"github.com/dsniels/market/core/types"
	"github.com/dsniels/market/internal/repo"
)

type ISucursal interface {
	GetAll(ctx context.Context) (*[]types.Sucursal, error)
	GetByID(ctx context.Context, id uint) (*types.Sucursal, error)
	Create(ctx context.Context, sucursal *types.Sucursal) error
	Update(ctx context.Context, id uint, sucursal *types.Sucursal) error
	DeleteByID(ctx context.Context, id uint) error
}

type Sucursal struct {
	repo repo.ISucursal
}

func (s *Sucursal) GetAll(ctx context.Context) (*[]types.Sucursal, error) {
	sucursals, err := s.repo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return sucursals, nil

}

func (s *Sucursal) GetByID(ctx context.Context, id uint) (*types.Sucursal, error) {
	sucursal, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return sucursal, nil
}

func (s *Sucursal) Create(ctx context.Context, sucursal *types.Sucursal) error {
	err := s.repo.Create(ctx, sucursal)
	if err != nil {
		return err
	}
	return err
}

func (s *Sucursal) Update(ctx context.Context, id uint, sucursal *types.Sucursal) error {
	cur, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	sucursal.ID = cur.ID
	_, err = s.repo.Update(sucursal)
	return err

}

func (s *Sucursal) DeleteByID(ctx context.Context, id uint) error {
	_, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func NewSucursal(repo repo.ISucursal) *Sucursal {
	return &Sucursal{
		repo: repo,
	}
}
