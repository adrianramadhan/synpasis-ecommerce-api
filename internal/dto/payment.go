package dto

type PaymentRequest struct {
	OrderId uint64 `json:"order_id"`
	Method  string `json:"method"`
}
