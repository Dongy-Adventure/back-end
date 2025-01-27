package router

import (
	"fmt"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	g    *gin.Engine
	conf *config.AppConfig
}

func NewRouter(g *gin.Engine, conf *config.AppConfig) *Router {
	return &Router{g, conf}
}

func (r *Router) Run(mongoDB *mongo.Database) {

	r.g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	v1 := r.g.Group("/api/v1")

	r.AddSellerRouter(v1, mongoDB)
	err := r.g.Run(":" + r.conf.Port)
	if err != nil {
		panic(fmt.Sprintf("Failed to run the server : %v", err))
	}
}
