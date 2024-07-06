package model

type Order struct {
	Id         uint64  `gorm:"primary_key;auto_increment"`
	UserId     uint64  `gorm:"column:user_id;not null"`
	CartId     uint64  `gorm:"column:cart_id;not null"`
	TotalPrice float64 `gorm:"column:total_price;type:decimal(10,2);not null"`
	Status     string  `gorm:"column:status;size:20;not null"`
	Common     `gorm:"embedded"`
}