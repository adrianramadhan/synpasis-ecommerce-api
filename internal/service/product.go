package service

import (
	"context"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/dto"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/repository"
)

type ProductService interface {
	GetProductsByCategory(ctx context.Context, categoryID uint32) ([]dto.ProductResponse, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProduct(productRepository repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (p *productService) GetProductsByCategory(ctx context.Context, categoryID uint32) ([]dto.ProductResponse, error) {
	products, err := p.productRepository.FindProductByCategory(ctx, categoryID)
	if err != nil {
		return nil, err

	}

	if len(products) == 0 {
		return []dto.ProductResponse{}, nil // Mengembalikan array kosong, bukan error
	}

	var productResponses []dto.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, dto.ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		})
	}

	return productResponses, nil
}
