package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redooz/go-jwt/modules/auth/middlewares"
	userControllers "github.com/redooz/go-jwt/modules/users/controllers"
)

func InitUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/hello", middlewares.JwtMiddleware, userControllers.HelloUser)
	}
}
