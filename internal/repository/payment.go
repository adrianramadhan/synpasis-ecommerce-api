package repository

import (
	"context"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/model"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *model.Payment) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPayment(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (r *paymentRepository) Create(ctx context.Context, payment *model.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}
