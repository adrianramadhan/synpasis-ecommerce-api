package repository

import (
	"context"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	FindByID(ctx context.Context, id uint64) (*model.Order, error)
	Update(ctx context.Context, order *model.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrder(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Create(ctx context.Context, order *model.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepository) FindByID(ctx context.Context, id uint64) (*model.Order, error) {
	var order model.Order
	err := r.db.WithContext(ctx).First(&order, id).Error
	return &order, err
}

func (r *orderRepository) Update(ctx context.Context, order *model.Order) error {
    return r.db.WithContext(ctx).Save(order).Error
}