package model

type Payment struct {
	Id      uint64  `gorm:"primary_key;auto_increment"`
	OrderId uint64  `gorm:"column:order_id;not null"`
	Amount  float64 `gorm:"column:amount;type:decimal(10,2);not null"`
	Status  string  `gorm:"column:status;size:20;not null"`
	Method  string  `gorm:"column:method;size:50;not null"`
	Common  `gorm:"embedded"`
}