package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Router) AddSellerRouter(rg *gin.RouterGroup, mongoDB *mongo.Database) {

	sellerRouter := rg.Group("seller")

	sellerRouter.POST("/", sellerCont.CreateSeller)
	sellerRouter.GET("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), sellerCont.GetSellers)
	sellerRouter.GET("/:seller_id", sellerCont.GetSellerByID)
	sellerRouter.PUT("/:seller_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), sellerCont.UpdateSeller)
	sellerRouter.POST("/:seller_id/transaction", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), sellerCont.AddTransaction)
	sellerRouter.GET("/:seller_id/balance", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), sellerCont.GetSellerBalanceByID)

}
