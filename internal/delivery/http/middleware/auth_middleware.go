package middleware

import (
	"net/http"
	"strings"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/utils"
	"github.com/adrianramadhan/synpasis-ecommerce-api/pkg/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func IsAuthenticated() echo.MiddlewareFunc {
	return echojwt.JWT([]byte(config.JwtSecret()))
}

func SetUserID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := utils.GetUserIDFromToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid token")
		}
		c.Set("user_id", userID)
		return next(c)
	}
}
