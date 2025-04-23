package router

import (
	"fmt"

	docs "github.com/Dongy-s-Advanture/back-end/docs"
	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/Dongy-s-Advanture/back-end/pkg/redis"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	rd "github.com/redis/go-redis/v9"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

func (r *Router) Run(mongoDB *mongo.Database, redisDB *rd.Client, s3Client *s3.Client) {

	r.g.Use(middleware.CORS())
	if r.conf.App.Env == "production" {
		r.g.Use(middleware.RateLimiter())
	}

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

	redisAdapter := redis.NewGoRedisAdapter(redisDB)

	// setup
	r.deps = NewDependencies(mongoDB, redisAdapter, s3Client, r.conf)

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
