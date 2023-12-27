package routes

import (
	"github.com/gin-gonic/gin"
	authControllers "github.com/redooz/go-jwt/modules/auth/controllers"
	userControllers "github.com/redooz/go-jwt/modules/users/controllers"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	authControllers := authControllers.AuthController{}
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", authControllers.SignUp)
	}

	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/hello", userControllers.HelloUser)
	}

	return r
}
