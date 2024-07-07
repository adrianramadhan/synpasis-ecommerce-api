package service

import (
	"context"
	"errors"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/dto"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/model"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/repository"
)

type CartService interface {
	AddProductToCart(ctx context.Context, userId uint64, request dto.AddToCartRequest) error
}

type cartService struct {
	cartRepository    repository.CartRepository
	productRepository repository.ProductRepository
}

func NewCart(cartRepository repository.CartRepository, productRepository repository.ProductRepository) CartService {
	return &cartService{
		cartRepository:    cartRepository,
		productRepository: productRepository,
	}
}

func (s *cartService) AddProductToCart(ctx context.Context, userId uint64, request dto.AddToCartRequest) error {
	// check if product exists
	product, err := s.productRepository.FindProductById(ctx, request.ProductId)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	// Check if quantity is valid
	if request.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	// Find or create cart for user
	cart, err := s.cartRepository.FindOrCreateByUserId(ctx, userId)
	if err != nil {
		return err
	}

	// Add product to cart
	cartItem := &model.CartItem{
		CartId:    cart.Id,
		ProductId: product.Id,
		Quantity:  request.Quantity,
	}

	return s.cartRepository.AddItem(ctx, cartItem)
}
