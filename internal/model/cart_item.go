package model

type CartItem struct {
	Id        uint64  `gorm:"primary_key;auto_increment"`
	CartId    uint64  `gorm:"column:cart_id;not null"`
	ProductId uint64  `gorm:"column:product_id;not null"`
	Quantity  uint8   `gorm:"column:quantity;not null"`
	Product   Product `gorm:"foreignkey:ProductId;references:Id"`
	Common    `gorm:"embedded"`
}
