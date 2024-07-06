package rest

import (
	"github.com/adrianramadhan/synpasis-ecommerce-api/internal/delivery/http/router"
)

func StartRest() {
	router.NewRouter().Init()
}
