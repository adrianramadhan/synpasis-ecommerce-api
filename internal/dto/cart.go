package dto

type AddToCartRequest struct {
	ProductId uint64 `json:"product_id"`
	Quantity  uint8  `json:"quantity"`
}

type CartItemResponse struct {
	ProductId   uint64  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    uint8   `json:"quantity"`
	Price       float64 `json:"price"`
}
