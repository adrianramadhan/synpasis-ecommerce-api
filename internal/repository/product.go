package repository

import (
	"context"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProductByCategory(ctx context.Context, categoryID uint32) ([]model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) FindProductByCategory(ctx context.Context, categoryID uint32) ([]model.Product, error) {
	var products []model.Product
	err := r.db.WithContext(ctx).Where("category_id = ?", categoryID).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
