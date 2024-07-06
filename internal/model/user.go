package model

type User struct {
	Id       uint64 `gorm:"primary_key;auto_increment"`
	Username string `gorm:"column:username;size:255;not null;unique"`
	Email    string `gorm:"column:email;size:255;not null;unique"`
	Password string `gorm:"column:password;size:255;not null"`
	Common   `gorm:"embedded"`
}