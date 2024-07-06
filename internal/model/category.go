package model

type Category struct {
	Id     uint32 `gorm:"primary_key;auto_increment"`
	Name   string `gorm:"column:name;size:100;not null;unique"`
	Common `gorm:"embedded"`
}