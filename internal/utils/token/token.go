package token

import (
	"errors"
	"fmt"
	"log"
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

func GenerateToken(conf *config.Config, userID string, tokenType tokenmode.TokenType) (string, error) {

	var tokenLifespan int32
	switch tokenType {
	case tokenmode.TokenMode.ACCESS_TOKEN:
		tokenLifespan = conf.Auth.AccessTokenLifespanMinutes
	case tokenmode.TokenMode.REFRESH_TOKEN:
		tokenLifespan = conf.Auth.RefreshTokenLifespanMinutes
	default:
		return "", errors.New("token type is invalid")
	}

	claims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Minute * time.Duration(tokenLifespan)).Unix(),
		"userID": userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	switch tokenType {
	case tokenmode.TokenMode.ACCESS_TOKEN:
		return token.SignedString([]byte(conf.Auth.AccessTokenSecret))
	case tokenmode.TokenMode.REFRESH_TOKEN:
		return token.SignedString([]byte(conf.Auth.RefreshTokenSecret))
	default:
		return "", errors.New("token type is invalid")
	}
}

func ValidateToken(c *gin.Context, tokenType tokenmode.TokenType) (*jwt.Token, error) {

	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("error loading .env")
	}
	tokenString := extractToken(c)

	if tokenString == "" {
		return nil, errors.New("no token given")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		switch tokenType {
		case tokenmode.TokenMode.ACCESS_TOKEN:
			return []byte(conf.Auth.AccessTokenSecret), nil
		case tokenmode.TokenMode.REFRESH_TOKEN:
			return []byte(conf.Auth.RefreshTokenSecret), nil
		default:
			return "", errors.New("token type is invalid")
		}
	})

	if err != nil {
		return nil, err
	}

	return token, err
}

func ExtractID(token *jwt.Token) (string, error) {
	if cliams, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, exits := cliams["userID"].(string)
		if !exits {
			return "", errors.New("userID not found in token")
		}
		return userID, nil
	}
	return "", nil
}
