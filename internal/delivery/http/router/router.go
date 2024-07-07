package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/delivery/http/handler"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/delivery/http/middleware"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/repository"
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/service"
	"github.com/adrianramadhan/synpasis-ecommerce-api/pkg/config"
	"github.com/adrianramadhan/synpasis-ecommerce-api/pkg/database"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type Router struct {
	userHandler    *handler.UserHandler
	productHandler *handler.ProductHandler
}

func NewRouter() *Router {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Login and register customers
	userRepository := repository.NewUser(db)
	userService := service.NewUser(userRepository)
	userHandler := handler.Newuser(userService)

	// Customer can view productlist by product category
	productRepository := repository.NewProduct(db)
	productService := service.NewProduct(productRepository)
	productHandler := handler.NewProduct(productService)

	return &Router{
		userHandler:    userHandler,
		productHandler: productHandler,
	}
}

func (r *Router) Init() {
	e := echo.New()
	e.Use(echoMiddleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/auth/register", r.userHandler.Register)
	e.POST("/auth/login", r.userHandler.Login)

	productGroup := e.Group("/products", middleware.IsAuthenticated())
	// productGroup.GET("", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	productGroup.GET("/category", r.productHandler.GetProductsByCategory)

	cartGroup := e.Group("/cart", middleware.IsAuthenticated())
	cartGroup.POST("/items", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	cartGroup.DELETE("/items{id}", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	cartGroup.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	orderGroup := e.Group("/orders", middleware.IsAuthenticated())
	orderGroup.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	orderGroup.GET("/{id}", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	paymentsGroup := e.Group("/payments", middleware.IsAuthenticated())
	paymentsGroup.POST("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	paymentsGroup.GET("/{id}", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.AppPort())))
}
