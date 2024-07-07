package service

import (
	"context"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/dto"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/model"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/repository"
)

type OrderService interface {
	CreateOrder(ctx context.Context, userId uint64) (*dto.OrderResponse, error)
}

type orderService struct {
	orderRepository repository.OrderRepository
	cartRepository  repository.CartRepository
}

func NewOrder(orderRepo repository.OrderRepository, cartRepo repository.CartRepository) OrderService {
	return &orderService{
		orderRepository: orderRepo,
		cartRepository:  cartRepo,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, userId uint64) (*dto.OrderResponse, error) {
	cart, err := s.cartRepository.FindOrCreateByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	cartItems, err := s.cartRepository.GetCartItems(ctx, cart.Id)
	if err != nil {
		return nil, err
	}

	var totalPrice float64
	for _, item := range cartItems {
		totalPrice += float64(item.Quantity) * item.Product.Price
	}

	order := &model.Order{
		UserId:     userId,
		CartId:     cart.Id,
		TotalPrice: totalPrice,
		Status:     "pending",
	}

	err = s.orderRepository.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	return &dto.OrderResponse{
		Id:         order.Id,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}, nil
}
