package repository

import (
	"context"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, input model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, input model.User) error {
	err := r.db.WithContext(ctx).Model(&model.User{}).Create(&input).Error
	if err != nil {
		return err
	}

	return nil
}
