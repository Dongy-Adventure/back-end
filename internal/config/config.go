package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type PaymentConfig struct {
	Public  string
	Private string
}

type AppConfig struct {
	Port string
	Env  string
}

type DbConfig struct {
	MongoURL      string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

type AuthConfig struct {
	AccessTokenSecret           string
	RefreshTokenSecret          string
	AccessTokenLifespanMinutes  int32
	RefreshTokenLifespanMinutes int32
}

type Config struct {
	App     AppConfig
	Auth    AuthConfig
	Db      DbConfig
	Payment PaymentConfig
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
	accessTokenLifeSpan, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_MINUTE_LIFESPAN"))
	if err != nil {
		return nil, err
	}
	refreshTokenLifeSpan, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_MINUTE_LIFESPAN"))
	if err != nil {
		return nil, err
	}

	authConfig := AuthConfig{
		AccessTokenSecret:           os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:          os.Getenv("REFRESH_TOKEN_SECRET"),
		AccessTokenLifespanMinutes:  int32(accessTokenLifeSpan),
		RefreshTokenLifespanMinutes: int32(refreshTokenLifeSpan),
	}
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}
	dbConfig := DbConfig{
		MongoURL:      os.Getenv("MONGODB_URL"),
		RedisAddr:     os.Getenv("REDIS_URL"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       redisDB,
	}

	paymentConfig := PaymentConfig{
		Public:  os.Getenv("OMISE_PUBLIC_KEY"),
		Private: os.Getenv("OMISE_PRIVATE_KEY"),
	}

	return &Config{
		App:     appConfig,
		Auth:    authConfig,
		Db:      dbConfig,
		Payment: paymentConfig,
	}, nil
}
