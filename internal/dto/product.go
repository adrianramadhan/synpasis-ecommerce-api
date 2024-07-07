package dto

type ProductResponse struct {
	Id          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       uint8   `json:"stock"`
}
