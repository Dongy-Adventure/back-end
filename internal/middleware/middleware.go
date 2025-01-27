package middleware

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleWare(tokenType tokenmode.TokenType) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
