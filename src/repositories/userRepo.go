package repositories

import (
	"context"
	"lsplanner-go/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetByID(ctx context.Context, id int) (*models.User, error)
	Add(ctx context.Context, user *models.User) error
	UpdateByID(ctx context.Context, id int, user *models.User) error
	DeleteByID(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]models.User, error)
	GetUserAreaID(ctx context.Context, id int) (uint, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	if result := ur.db.First(&user, id); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *userRepo) Add(ctx context.Context, user *models.User) error {
	if result := ur.db.Create(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepo) UpdateByID(ctx context.Context, id int, user *models.User) error {
	if result := ur.db.Model(&models.User{}).Where("id = ?", id).Updates(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepo) DeleteByID(ctx context.Context, id int) error {
	if result := ur.db.Delete(&models.User{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepo) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	if result := ur.db.Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (ur *userRepo) GetUserAreaID(ctx context.Context, id int) (uint, error) {
	var user models.User
	if result := ur.db.First(&user, id); result.Error != nil {
		return 0, result.Error
	}
	return user.AreaID, nil
}
