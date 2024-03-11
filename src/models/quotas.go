package models

import (
	"time"

	"gorm.io/gorm"
)

type Quota struct {
	gorm.Model
	CurrentCount int
	MaxCount     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	User         User `gorm:"foreignKey:ID"` // This line references the User's primary key
}
