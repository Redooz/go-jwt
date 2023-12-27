package services

import (
	"github.com/redooz/go-jwt/database"
	"github.com/redooz/go-jwt/modules/auth/dtos"
	"github.com/redooz/go-jwt/modules/users/models"
)

type UserService struct {
	user models.User
}

func (u UserService) CreateUser(signUpDto *dtos.SignUpRequestDto) (userEmail string, err error) {
	u.user.Email = signUpDto.Email
	u.user.Password = signUpDto.Password

	result := database.DB.Create(&u.user)

	if result.Error != nil {
		return "", result.Error
	}

	return u.user.Email, nil
}

func (u UserService) GetUserByEmail(email string) (user *models.User, err error) {
	result := database.DB.Where(&models.User{Email: email}).First(&u.user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &u.user, nil
}

func (u UserService) GetUserById(id uint) (user *models.User, err error) {
	result := database.DB.Where(&models.User{ID: id}).First(&u.user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &u.user, nil
}