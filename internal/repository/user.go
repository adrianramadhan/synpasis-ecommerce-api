package repository

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(ctx context.Context) (string, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Login(ctx context.Context) (string, error) {
	panic("implement me");
}