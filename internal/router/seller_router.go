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
	repo := repository.NewSellerRepository(mongoDB, "sellers", "reviews")
	serv := service.NewSellerService(repo)
	cont := controller.NewSellerController(serv)

	sellerRouter := rg.Group("seller")

	sellerRouter.POST("/", cont.CreateSeller)
	sellerRouter.GET("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), cont.GetSellers)
	sellerRouter.GET("/:seller_id", cont.GetSellerByID)
	sellerRouter.PUT("/:seller_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), cont.UpdateSeller)
	sellerRouter.POST("/:seller_id/withdraw", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), cont.WithdrawSellerBalance)
	sellerRouter.GET("/:seller_id/balance", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), cont.GetSellerBalanceByID)

}
