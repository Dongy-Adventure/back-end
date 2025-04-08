package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r Router) AddBuyerRouter(rg *gin.RouterGroup) {

	buyerCont := r.deps.BuyerController

	buyerRouter := rg.Group("buyer")

	buyerRouter.POST("/", buyerCont.CreateBuyer)
	buyerRouter.GET("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis), buyerCont.GetBuyers)
	buyerRouter.GET("/:buyer_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis), buyerCont.GetBuyerByID)
	buyerRouter.PUT("/:buyer_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis), buyerCont.UpdateBuyer)
	buyerRouter.POST("/:buyer_id/cart", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis), buyerCont.UpdateProductInCart)
	buyerRouter.DELETE("/:buyer_id/cart/:product_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis), buyerCont.DeleteProductFromCart)

}
