package services

import (
	"context"
	"fmt"
)

type ISession interface {
	Login(ctx context.Context, email string, password string) (*string, error)
}

type Session struct {
	userSvc  IUser
	tokenSvc *Token
}

func (s *Session) Login(ctx context.Context, email string, password string) (*string, error) {

	user, err := s.userSvc.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if ok, _ := user.ValidatePassword(password); !ok {
		return nil, fmt.Errorf("incorrect password or email")
	}

	token, err := s.tokenSvc.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func NewSession(svc IUser, tokenSvc *Token) *Session {
	return &Session{
		tokenSvc: tokenSvc,
		userSvc:  svc,
	}

}
