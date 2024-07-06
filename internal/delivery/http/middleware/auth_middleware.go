package middleware

import (
	"github.com/adrianramadhan/synpasis-ecommerce-api/pkg/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func 'IsAuthenticated() echo.MiddlewareFunc {
	return echojwt.JWT([]byte(config.JwtSecret()))
}
