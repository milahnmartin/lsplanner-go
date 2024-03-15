package repositories

import (
	"context"
	"errors"
	"lsplanner-go/models"

	"gorm.io/gorm"
)

type QuotaRepo interface {
	GetByID(ctx context.Context, id int) (*models.Quota, error)
	InitQuota(ctx context.Context, id int, limit int) (*models.Quota, error)
}

type quotaRepo struct {
	db *gorm.DB
}

func NewQuotaRepo(db *gorm.DB) QuotaRepo {
	return &quotaRepo{db: db}
}

func (ur *quotaRepo) GetByID(ctx context.Context, id int) (*models.Quota, error) {
	var quota models.Quota
	if result := ur.db.First(&quota, id); result.Error != nil {
		return nil, result.Error
	}
	return &quota, nil
}

func (ur *quotaRepo) InitQuota(ctx context.Context, id int, limit int) (*models.Quota, error) {
	var quota models.Quota

	// Check if a quota record with the given user ID already exists
	if err := ur.db.Where("user_id = ?", id).First(&quota).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err // Return error if it's not "record not found" error
		}
	} else {
		// If a quota record already exists for the user, return an error
		return nil, errors.New("user already has a quota record")
	}

	// If no quota record exists, create a new one
	quota = models.Quota{
		MaxCount:     limit,
		CurrentCount: 0,
		UserID:       id,
	}

	if err := ur.db.Create(&quota).Error; err != nil {
		return nil, err
	}

	return &quota, nil
}
