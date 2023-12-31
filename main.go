package main

import (
	"github.com/redooz/go-jwt/config"
	"github.com/redooz/go-jwt/database"
	"github.com/redooz/go-jwt/docs"
	"github.com/redooz/go-jwt/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.LoadEnvVariables()
	database.ConnectToDb()
}

func main() {
	r := routes.InitRoutes()

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(string(":" + config.PORT))
}
