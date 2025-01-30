package token

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func extractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
	}
	return ""

}

func GenerateToken(tokenType tokenmode.TokenType) (string, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading .env File")
	}
	var token_lifespan int
	switch tokenType {
	case tokenmode.TokenMode.ACCESS_TOKEN:
		token_lifespan, err = strconv.Atoi(config.Auth.AccessTokenLifespanMinutes)
	case tokenmode.TokenMode.REFRESH_TOKEN:
		token_lifespan, err = strconv.Atoi(config.Auth.RefreshTokenLifespanMinutes)
	default:
		return "", errors.New("token type is invalid")
	}
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	switch tokenType {
	case tokenmode.TokenMode.ACCESS_TOKEN:
		return token.SignedString([]byte(config.Auth.AccessTokenSecret))
	case tokenmode.TokenMode.REFRESH_TOKEN:
		return token.SignedString([]byte(config.Auth.RefreshTokenSecret))
	default:
		return "", errors.New("token type is invalid")
	}
}

func ValidateToken(c *gin.Context, tokenType tokenmode.TokenType) error {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("Error loading .env file in validate token")
	}
	tokenString := extractToken(c)

	if tokenString == "" {
		return errors.New("no token given")
	}

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		switch tokenType {
		case tokenmode.TokenMode.ACCESS_TOKEN:
			return []byte(config.Auth.AccessTokenSecret), nil
		case tokenmode.TokenMode.REFRESH_TOKEN:
			return []byte(config.Auth.RefreshTokenSecret), nil
		default:
			return "", errors.New("token type is invalid")
		}
	})

	if err != nil {
		return err
	}

	return nil
}
