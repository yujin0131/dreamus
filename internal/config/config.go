package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		logrus.WithError(err).Warn("Failed to load .env file")
	}
}
