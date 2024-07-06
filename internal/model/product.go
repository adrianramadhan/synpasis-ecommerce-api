package model

type Product struct {
	Id          uint64  `gorm:"primary_key;auto_increment"`
	Name        string  `gorm:"column:name;size:255;not null"`
	Description string  `gorm:"column:description;type:text"`
	Price       float64 `gorm:"column:price;type:decimal(10,2);not null"`
	Stock       uint8   `gorm:"column:stock;not null"`
	CategoryId  uint32  `gorm:"column:category_id"`
	Common      `gorm:"embedded"`
}