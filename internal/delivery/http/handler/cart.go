package handler

import (
	"net/http"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/dto"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/service"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	service service.CartService
}

func NewCart(service service.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) AddToCart(c echo.Context) error {
	var request dto.AddToCartRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}

	// Assume we have middleware that sets user ID in context
	userID := c.Get("user_id").(uint64)

	err := h.service.AddProductToCart(c.Request().Context(), userID, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Product added to cart successfully")
}
