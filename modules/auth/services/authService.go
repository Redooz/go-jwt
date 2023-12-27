package services

import (
	"net/http"

	"github.com/redooz/go-jwt/database"
	"github.com/redooz/go-jwt/modules/auth/dtos"
	"github.com/redooz/go-jwt/modules/users/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	user models.User
}

func (a AuthService) SignUp(signUpDto *dtos.SignUpRequestDto) (response *dtos.SignUpResponseDto, httpStatus int, err error) {
	a.user.Email = signUpDto.Email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpDto.Password), bcrypt.DefaultCost)

	if err != nil {
		// return error 500
		return nil, http.StatusInternalServerError, err
	}

	a.user.Password = string(hashedPassword)

	result := database.DB.Create(&a.user)

	if result.Error != nil {
		// return error 500
		return nil, http.StatusInternalServerError, result.Error
	}

	// return success
	return &dtos.SignUpResponseDto{Email: a.user.Email}, http.StatusOK, nil
}

func (a AuthService) SignIn(signInDto *dtos.SignInRequestDto) (response *dtos.SignInResponseDto, httpStatus int, err error) {
	return nil, 0, nil
}