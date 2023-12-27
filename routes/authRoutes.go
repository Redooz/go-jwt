package routes

import (
    "github.com/gin-gonic/gin"
    authControllers "github.com/redooz/go-jwt/modules/auth/controllers"
)

func InitAuthRoutes(r *gin.Engine) {
    authController := authControllers.AuthController{}
    authRoutes := r.Group("/auth")
    {
        authRoutes.POST("/signup", authController.SignUp)
        authRoutes.POST("/signin", authController.SignIn)
    }
}