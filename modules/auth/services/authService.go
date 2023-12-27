package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redooz/go-jwt/config"
	"github.com/redooz/go-jwt/modules/auth/dtos"
	userServices "github.com/redooz/go-jwt/modules/users/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserService userServices.UserService
}

func (a AuthService) SignUp(signUpDto *dtos.SignUpRequestDto) (response *dtos.SignUpResponseDto, httpStatus int, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpDto.Password), bcrypt.DefaultCost)

	if err != nil {
		// return error 500
		return nil, http.StatusInternalServerError, err
	}

	signUpDto.Password = string(hashedPassword)

	// create user
	email, err := a.UserService.CreateUser(signUpDto)

	if err != nil {
		// return error 500
		return nil, http.StatusInternalServerError, err
	}

	// return success
	return &dtos.SignUpResponseDto{Email: email}, http.StatusOK, nil
}

func (a AuthService) SignIn(signInDto *dtos.SignInRequestDto) (response *dtos.SignInResponseDto, httpStatus int, err error) {
	user, err := a.UserService.GetUserByEmail(signInDto.Email)

	if err != nil {
		customError := errors.New("invalid email or password")
		// return not found
		return nil, http.StatusNotFound, customError
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signInDto.Password))

	if err != nil {
		customError := errors.New("invalid email or password")
		// return unauthorized with custom message error
		return nil, http.StatusUnauthorized, customError
	}

	// generate jwt token
	token, err := generateJwtToken(user.ID)

	if err != nil {
		// return error 500
		return nil, http.StatusInternalServerError, err
	}

	// return success
	return &dtos.SignInResponseDto{Email: user.Email, Token: token}, http.StatusOK, nil
}

func generateJwtToken(userId uint) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err = token.SignedString([]byte(config.SECRET))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
