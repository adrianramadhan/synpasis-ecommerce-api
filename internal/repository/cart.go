package repository

import (
	"context"
	"errors"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/model"
	"gorm.io/gorm"
)

type CartRepository interface {
	FindOrCreateByUserId(ctx context.Context, userId uint64) (*model.Cart, error)
	AddItem(ctx context.Context, cartItem *model.CartItem) error
	GetCartItems(ctx context.Context, cartId uint64) ([]model.CartItem, error)
	DeleteCartItem(ctx context.Context, cartId uint64, productId uint64) error
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

func (r *cartRepository) GetCartItems(ctx context.Context, cartId uint64) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	err := r.db.WithContext(ctx).Where("cart_id = ?", cartId).Preload("Product").Find(&cartItems).Error
	return cartItems, err
}

func (r *cartRepository) DeleteCartItem(ctx context.Context, cartId uint64, productId uint64) error {
	result := r.db.WithContext(ctx).Where("cart_id = ? AND product_id = ?", cartId, productId).Delete(&model.CartItem{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("cart item not found")
	}
	return nil
}
