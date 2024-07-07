package service

import (
	"context"
	"errors"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/dto"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/model"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/repository"
)

type PaymentService interface {
	ProcessPayment(ctx context.Context, userId uint64, request dto.PaymentRequest) error
}

type paymentService struct {
	paymentRepository repository.PaymentRepository
	orderRepository   repository.OrderRepository
}

func NewPayment(paymentRepo repository.PaymentRepository, orderRepo repository.OrderRepository) PaymentService {
	return &paymentService{
		paymentRepository: paymentRepo,
		orderRepository:   orderRepo,
	}
}

func (s *paymentService) ProcessPayment(ctx context.Context, userId uint64, request dto.PaymentRequest) error {
	order, err := s.orderRepository.FindByID(ctx, request.OrderId)
	if err != nil {
		return err
	}

	if order.UserId != userId {
		return errors.New("unauthorized")
	}

	payment := &model.Payment{
		OrderId: order.Id,
		Amount:  order.TotalPrice,
		Method:  request.Method,
		Status:  "completed",
	}

	err = s.paymentRepository.Create(ctx, payment)
	if err != nil {
		return err
	}

	order.Status = "paid"
	return s.orderRepository.Update(ctx, order)
}
