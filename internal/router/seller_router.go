package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/controller"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Router) AddSellerRouter(rg *gin.RouterGroup, mongoDB *mongo.Database) {
	repo := repository.NewSellerRepository(mongoDB, "sellers")
	serv := service.NewSellerService(repo)
	cont := controller.NewSellerController(serv)

	sellerRouter := rg.Group("seller")

	sellerRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.CreateSeller)
	sellerRouter.GET("/:seller_id", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.GetSellerByID)
	sellerRouter.PUT("/:seller_id", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.GetSellerByID)

}
