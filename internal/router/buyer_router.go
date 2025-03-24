package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Router) AddBuyerRouter(rg *gin.RouterGroup, mongoDB *mongo.Database) {

	buyerRouter := rg.Group("buyer")

	buyerRouter.POST("/", buyerCont.CreateBuyer)
	buyerRouter.GET("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), buyerCont.GetBuyers)
	buyerRouter.GET("/:buyer_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), buyerCont.GetBuyerByID)
	buyerRouter.PUT("/:buyer_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), buyerCont.UpdateBuyer)
	buyerRouter.PATCH("/:buyer_id/cart", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), buyerCont.UpdateProductInCart)
  buyerRouter.DELETE("/:buyer_id/cart/:product_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), buyerCont.DeleteProductFromCart)

}
