package middleware

import (
	"net/http"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/pkg/redis"
	token "github.com/Dongy-s-Advanture/back-end/pkg/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter"
	memory "github.com/ulule/limiter/drivers/store/memory"
)

func JWTAuthMiddleWare(tokenType int, redisClient redis.IRedisClient, conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tkn, err := token.ValidateToken(c, conf, redisClient, tokenType)
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

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, PATCH, PUT, GET, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
func RateLimiter() gin.HandlerFunc {

	rate := limiter.Rate{
		Period: time.Minute,
		Limit:  100,
	}

	store := memory.NewStore()
	limiter := limiter.New(store, rate)

	return func(c *gin.Context) {

		ip := c.ClientIP()
		context, err := limiter.Get(c, ip)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
			return
		}

		if context.Reached {
			c.AbortWithStatusJSON(429, gin.H{"error": "too many requests"})
			return
		}

		c.Next()
	}
}
