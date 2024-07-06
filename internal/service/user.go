package service

import "context"

type UserService interface {
	Login(ctx context.Context) (string, error)
}

type userService struct {
}

func NewUser() UserService {
	return &userService{}
}

func (u *userService) Login(ctx context.Context) (string, error) {
	panic("implement me")
}