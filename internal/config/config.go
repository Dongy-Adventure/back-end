package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port string
	Env  string
}

type DbConfig struct {
	MongoURL string
}

type AuthConfig struct {
	AccessTokenSecret           string
	RefreshTokenSecret          string
	AccessTokenLifespanMinutes  string
	RefreshTokenLifespanMinutes string
}

type Config struct {
	App  AppConfig
	Auth AuthConfig
	Db   DbConfig
}

func LoadConfig() (*Config, error) {
	dir, err := os.Getwd() // Capture both the directory and the error
	if err != nil {
		fmt.Println("Error getting working directory:", err)
	}
	fmt.Println("Current Working Directory:", dir)
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	appConfig := AppConfig{
		Env:  os.Getenv("APP_ENV"),
		Port: os.Getenv("APP_PORT"),
	}

	authConfig := AuthConfig{
		AccessTokenSecret:           os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:          os.Getenv("REFRESH_TOKEN_SECRET"),
		AccessTokenLifespanMinutes:  os.Getenv("ACCESS_TOKEN_MINUTE_LIFESPAN"),
		RefreshTokenLifespanMinutes: os.Getenv("REFRESH_TOKEN_MINUTE_LIFESPAN"),
	}

	dbConfig := DbConfig{
		MongoURL: os.Getenv("MONGODB_URL"),
	}

	return &Config{
		App:  appConfig,
		Auth: authConfig,
		Db:   dbConfig,
	}, nil
}
