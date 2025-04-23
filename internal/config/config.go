package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Mongo MongoConfig `envPrefix:"MONGO_"`
}

type MongoConfig struct {
	URI      string `env:"URI"`
	Database string `env:"DATABASE"`
}

// @WireSet("Config")
func NewConfig() *Config {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found or error loading it. Falling back to system environment variables.")
	}

	config := &Config{}

	// Parse environment variables into the config struct
	if err := env.Parse(config); err != nil {
		logrus.Error("Failed to parse environment variables into Config struct:", err)
	}

	return config
}
