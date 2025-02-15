package middleware

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	token "github.com/Dongy-s-Advanture/back-end/internal/utils/token"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleWare(tokenType int) gin.HandlerFunc {
	return func(c *gin.Context) {
		tkn, err := token.ValidateToken(c, tokenType)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Success: false, Status: http.StatusUnauthorized, Error: "Unauthorized", Message: err.Error()})
			c.Abort()
			return
		}
		userID, err := token.ExtractID(tkn)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Success: false, Status: http.StatusUnauthorized, Error: "No userID in token", Message: err.Error()})
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
