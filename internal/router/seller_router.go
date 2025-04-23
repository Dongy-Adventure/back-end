package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r Router) AddSellerRouter(rg *gin.RouterGroup) {

	sellerCont := r.deps.SellerController
	sellerRouter := rg.Group("seller")

	sellerRouter.POST("/", sellerCont.CreateSeller)
	sellerRouter.GET("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), sellerCont.GetSellers)
	sellerRouter.GET("/:seller_id", sellerCont.GetSellerByID)
	sellerRouter.PUT("/:seller_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), sellerCont.UpdateSeller)
	sellerRouter.POST("/:seller_id/withdraw", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), sellerCont.WithdrawSellerBalance)
	sellerRouter.GET("/:seller_id/balance", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), sellerCont.GetSellerBalanceByID)
}
