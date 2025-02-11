package router

import (
	"fmt"

	docs "github.com/Dongy-s-Advanture/back-end/docs"
	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	g    *gin.Engine
	conf *config.Config
}

func NewRouter(g *gin.Engine, conf *config.Config) *Router {
	return &Router{g, conf}
}

func (r *Router) Run(mongoDB *mongo.Database) {

	// CORS setting
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
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

	// Add related path
	r.AddSellerRouter(v1, mongoDB)
	r.AddBuyerRouter(v1, mongoDB)
	r.AddAuthRouter(v1, mongoDB)
	r.AddProductRouter(v1, mongoDB)
	err := r.g.Run(":" + r.conf.App.Port)
	if err != nil {
		panic(fmt.Sprintf("Failed to run the server : %v", err))
	}
}
