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
	cartHandler    *handler.CartHandler
	orderHandler   *handler.OrderHandler
	paymentHandler *handler.PaymentHandler
}

func NewRouter() *Router {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepository := repository.NewUser(db)
	userService := service.NewUser(userRepository)
	userHandler := handler.Newuser(userService)

	productRepository := repository.NewProduct(db)
	productService := service.NewProduct(productRepository)
	productHandler := handler.NewProduct(productService)

	cartRepository := repository.NewCart(db)
	cartService := service.NewCart(cartRepository, productRepository)
	cartHandler := handler.NewCart(cartService)

	orderRepository := repository.NewOrder(db)
	orderService := service.NewOrder(orderRepository, cartRepository)
	orderHandler := handler.NewOrder(orderService)

	paymentRepository := repository.NewPayment(db)
	paymentService := service.NewPayment(paymentRepository, orderRepository)
	paymentHandler := handler.NewPayment(paymentService)

	return &Router{
		userHandler:    userHandler,
		productHandler: productHandler,
		cartHandler:    cartHandler,
		orderHandler:   orderHandler,
		paymentHandler: paymentHandler,
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
	productGroup.GET("/category", r.productHandler.GetProductsByCategory)

	cartGroup := e.Group("/cart", middleware.IsAuthenticated(), middleware.SetUserID)
	cartGroup.POST("/add", r.cartHandler.AddToCart)
	cartGroup.GET("/items", r.cartHandler.GetCartItems)
	cartGroup.DELETE("/delete", r.cartHandler.DeleteFromCart)

	orderGroup := e.Group("/orders", middleware.IsAuthenticated(), middleware.SetUserID)
	orderGroup.POST("/create", r.orderHandler.CreateOrder)

	paymentGroup := e.Group("/payments", middleware.IsAuthenticated(), middleware.SetUserID)
	paymentGroup.POST("/process", r.paymentHandler.ProcessPayment)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.AppPort())))
}
