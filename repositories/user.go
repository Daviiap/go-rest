package repositories

import (
	"go_rest_api/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) GetUserByEmail(email string) models.User {
	var user models.User

	repo.db.Where("email = ?", email).First(&user)

	return user
}
