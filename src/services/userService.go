package services

import (
	"context"
	"lsplanner-go/models"
	"lsplanner-go/repositories"
)

func CreateUser(ctx context.Context, user *models.User, userRepo repositories.UserRepo) error {
	return userRepo.Add(ctx, user)
}

func GetAllUsers(ctx context.Context, userRepo repositories.UserRepo) ([]models.User, error) {
	return userRepo.GetAll(ctx)
}

func GetUserByID(ctx context.Context, id int, userRepo repositories.UserRepo) (*models.User, error) {
	return userRepo.GetByID(ctx, id)
}

func UpdateUserByID(ctx context.Context, id int, user *models.User, userRepo repositories.UserRepo) error {
	return userRepo.UpdateByID(ctx, id, user)
}

func DeleteUserByID(ctx context.Context, id int, userRepo repositories.UserRepo) error {
	return userRepo.DeleteByID(ctx, id)
}
