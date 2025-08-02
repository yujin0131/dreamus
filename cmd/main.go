// @title Dreamus Build API
// @version 1.0
// @description API for managing build metadata (commit, Docker tag, status)
// @host localhost:8080
// @BasePath /
// @schemes http
package main

import (
	_ "dreamus/docs"

	"dreamus/internal/config"
	"dreamus/internal/db"
	"dreamus/internal/models"
	"dreamus/internal/routes"
	"os"

	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.LoadEnv()
	db.ConnectDB()

	if err := db.DB.AutoMigrate(&models.Build{}); err != nil {
		logrus.WithError(err).Warn("Failed AutoMigrate")
	}
	r := routes.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		logrus.WithError(err).Fatal("Failed to run server")
	}
}
