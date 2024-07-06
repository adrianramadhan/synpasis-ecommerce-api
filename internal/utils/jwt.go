package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/adrianramadhan/synpasis-ecommerce-api/pkg/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtSecret = config.JwtSecret()

type JWTCustomClaim struct {
	ID uint64 `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(id uint64) (string, error) {
	claims := JWTCustomClaim{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Local().Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func ExtractToken(c echo.Context) (*JWTCustomClaim, error) {
	tokenFromHeader := c.Request().Header.Get("Authorization")

	sanitizedTokenBearer := strings.Replace(tokenFromHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(sanitizedTokenBearer, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)

		userId := claims["id"].(uint64)

		customClaims := &JWTCustomClaim{
			ID: userId,
		}

		return customClaims, nil
	}

	return nil, errors.New("invalid token")
}
