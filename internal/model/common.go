package model

import "time"

type Common struct {
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at;not null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at;not null"`
}
