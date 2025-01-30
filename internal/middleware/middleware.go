package middleware

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/utils/token"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleWare(tokenType tokenmode.TokenType) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.ValidateToken(c, tokenType)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Success: false, Status: http.StatusUnauthorized, Error: "Unauthorized", Message: err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
