package handler

import (
	"net/http"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/dto"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/service"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPayment(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) ProcessPayment(c echo.Context) error {
	var request dto.PaymentRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	userID := c.Get("user_id").(uint64)

	err := h.service.ProcessPayment(c.Request().Context(), userID, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Payment processed successfully")
}
