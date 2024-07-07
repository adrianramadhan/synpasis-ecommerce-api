package dto

type AddToCartRequest struct {
	ProductId uint64 `json:"product_id"`
	Quantity  uint8  `json:"quantity"`
}
