package handler

import (
	"net/http"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/dto"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/service"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service service.UserService
}

func Newuser(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(c echo.Context) error {
	var request dto.UserRegisterRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := h.service.Register(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "success")

}

func (h *UserHandler) Login(c echo.Context) error {
	var request dto.UserLoginRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := h.service.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}
