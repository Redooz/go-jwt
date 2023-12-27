package main

import (
	"github.com/redooz/go-jwt/config"
	"github.com/redooz/go-jwt/database"
	"github.com/redooz/go-jwt/routes"
)

func init() {
	config.LoadEnvVariables()
	database.ConnectToDb()
}

func main() {
	r := routes.InitRoutes()

	r.Run(string(":" + config.PORT))
}
