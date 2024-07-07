package handler

import (
	"net/http"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/service"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrder(service service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	userID := c.Get("user_id").(uint64)

	order, err := h.service.CreateOrder(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, order)
}
