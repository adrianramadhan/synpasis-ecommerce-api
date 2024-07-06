package model

type Common struct {
	CreatedAt string `gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt string `gorm:"autoUpdateTime;column:updated_at;not null"`
}