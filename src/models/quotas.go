package models

import (
	"time"

	"gorm.io/gorm"
)

type Quota struct {
	gorm.Model
	CurrentCount int            `gorm:"default:0" json:"current_count"`
	MaxCount     int            `gorm:"default:0" json:"max_count"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UserID       int            `json:"user_id"`
}
