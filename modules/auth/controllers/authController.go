package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redooz/go-jwt/modules/auth/dtos"
	"github.com/redooz/go-jwt/modules/auth/services"
)

type AuthController struct {
	service services.AuthService
}

func (controller AuthController) SignUp(c *gin.Context) {
	var body dtos.SignUpRequestDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(body)
	validationErrors := err.(validator.ValidationErrors)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Error()})
		return
	}

	response, httpStatus, err := controller.service.SignUp(&body)

	if err != nil {
		c.JSON(httpStatus, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(httpStatus, gin.H{
		"message": "User Created",
		"data":    response,
	})
}
