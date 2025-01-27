package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/controller"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Router) AddAuthRouter(rg *gin.RouterGroup, mongoDB *mongo.Database) {
	sellerRepo := repository.NewSellerRepository(mongoDB, "sellers")
	serv := service.NewAuthService(sellerRepo)
	cont := controller.NewAuthController(serv)

	sellerRouter := rg.Group("auth")

	sellerRouter.POST("/", cont.SellerLogin)

}
