package model

type Cart struct {
	Id     uint64 `gorm:"primary_key;auto_increment"`
	UserId uint64 `gorm:"column:user_id;not null"`
	Status string `gorm:"column:status;size:20;not null"`
	Common `gorm:"embedded"`
}