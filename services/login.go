package services

import (
	"errors"
	"go_rest_api/models"
	"go_rest_api/repositories"
	"go_rest_api/utils"
)

type LoginService struct {
	userRepo *repositories.UserRepo
}

func NewLoginService(userRepo *repositories.UserRepo) *LoginService {
	return &LoginService{
		userRepo: userRepo,
	}
}

func (service *LoginService) Login(email, password string) (string, models.User, error) {
	user := service.userRepo.GetUserByEmail(email)

	if user.Name == "" {
		return "", user, errors.New("user not found")
	}

	if !user.VerifyPassword(password) {
		return "", user, errors.New("invalid password")
	}

	token, err := utils.GenerateToken(user.ID)

	return token, user, err
}
