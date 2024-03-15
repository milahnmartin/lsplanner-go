package services

import (
	"context"
	"lsplanner-go/models"
	"lsplanner-go/repositories"
)

func GetByID(ctx context.Context, id int, quotaRepo repositories.QuotaRepo) (*models.Quota, error) {
	return quotaRepo.GetByID(ctx, id)
}

func InitQuota(ctx context.Context, id int, limit int, quotaRepo repositories.QuotaRepo) (*models.Quota, error) {
	return quotaRepo.InitQuota(ctx, id, limit)
}
