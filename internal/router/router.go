package router

import (
	"fmt"

	"time"

	docs "github.com/Dongy-s-Advanture/back-end/docs"
	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	g    *gin.Engine
	conf *config.Config
	deps *Dependencies
}

func NewRouter(g *gin.Engine, conf *config.Config) *Router {
	return &Router{g, conf, nil}
}

func (r *Router) Run(mongoDB *mongo.Database, redisDB *redis.Client) {

	// CORS setting
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowMethods = []string{"OPTIONS", "PATCH", "PUT", "GET", "POST", "DELETE"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"} // Allow Authorization header
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true // If you are using cookies or Authorization header

	// Optional: Handle preflight cache
	corsConfig.MaxAge = 12 * time.Hour

	r.g.Use(cors.New(corsConfig))

	r.g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Swagger setting
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// versioning
	v1 := r.g.Group("/api/v1")

	// setup
	r.deps = NewDependencies(mongoDB, redisDB, r.conf)

	// Add related path
	r.AddSellerRouter(v1)
	r.AddBuyerRouter(v1)
	r.AddAuthRouter(v1)
	r.AddProductRouter(v1)
	r.AddOrderRouter(v1)
	r.AddReviewRouter(v1)
	r.AddAppointmentRouter(v1)
	r.AddPaymentRouter(v1)
	r.AddAdvertisementRouter(v1)

	err := r.g.Run(":" + r.conf.App.Port)
	if err != nil {
		panic(fmt.Sprintf("Failed to run the server : %v", err))
	}
}
