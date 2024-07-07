package repository

import (
	"context"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/model"
	"gorm.io/gorm"
)

type CartRepository interface {
	FindOrCreateByUserId(ctx context.Context, userId uint64) (*model.Cart, error)
	AddItem(ctx context.Context, cartItem *model.CartItem) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCart(db *gorm.DB) CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r *cartRepository) FindOrCreateByUserId(ctx context.Context, userId uint64) (*model.Cart, error) {
	var cart model.Cart
	err := r.db.WithContext(ctx).FirstOrCreate(&cart, model.Cart{UserId: userId, Status: "active"}).Error
	return &cart, err
}

func (r *cartRepository) AddItem(ctx context.Context, cartItem *model.CartItem) error {
	return r.db.WithContext(ctx).Create(cartItem).Error
}
