package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/service"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProduct(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) GetProductsByCategory(c echo.Context) error {
	categoryIDStr := c.QueryParam("category_id")
	log.Println("Received category_id: ", categoryIDStr)

	if categoryIDStr == "" {
		return c.JSON(http.StatusBadRequest, "Category ID is required")
	}

	categoryId, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid category ID: must be a positive integer")
	}

	products, err := h.service.GetProductsByCategory(c.Request().Context(), uint32(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusOK, products)
}
