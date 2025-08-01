package main

import (
	"dreamus/internal/config"
	"dreamus/internal/db"
	"dreamus/internal/models"
	"dreamus/internal/routes"

	"github.com/sirupsen/logrus"
)

func main() {
	config.LoadEnv()
	db.ConnectDB()

	if err := db.DB.AutoMigrate(&models.Build{}); err != nil {
		logrus.WithError(err).Warn("Failed AutoMigrate")
	}
	r := routes.SetupRouter()
	r.Run(":8080")
}
