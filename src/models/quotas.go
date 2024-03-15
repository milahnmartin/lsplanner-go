package models

import (
	"time"

	"gorm.io/gorm"
)

type Quota struct {
	gorm.Model
	CurrentCount int `gorm:"default:0" json:"current_count"`
	MaxCount     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
