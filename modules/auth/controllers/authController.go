package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redooz/go-jwt/modules/auth/dtos"
	"github.com/redooz/go-jwt/modules/auth/services"
)

type AuthController struct {
	service services.AuthService
}

func (controller AuthController) SignUp(c *gin.Context) {
	var body dtos.SignUpRequestDto

	log.Default().Println("Sign Up")

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Request",
		})
	}

	response, httpStatus, err := controller.service.SignUp(&body)

	if err != nil {
		c.JSON(httpStatus, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(httpStatus, gin.H{
		"message": "User Created",
		"data":    response,
	})
}
