package dto

type OrderResponse struct {
	Id         uint64  `json:"id"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}
