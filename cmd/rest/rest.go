package rest

import (
	"fmt"
	"net/http"

	"github.com/adrianramadhan/synpasis-ecommerce-api/pkg/config"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartRest() {
	e := echo.New()
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.AppPort())))
}