package models

import "gorm.io/gorm"

type LoadsheddingArea struct {
	gorm.Model
	Schedule string `jsonb:"schedule"`
}
