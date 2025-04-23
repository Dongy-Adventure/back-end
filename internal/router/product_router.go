package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r Router) AddProductRouter(rg *gin.RouterGroup) {

	productCont := r.deps.ProductController
	productRouter := rg.Group("product")

	productRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), productCont.CreateProduct)
	productRouter.GET("/", productCont.GetProducts)
	productRouter.GET("/:product_id", productCont.GetProductByID)
	productRouter.GET("/seller/:seller_id", productCont.GetProductsBySellerID)
	productRouter.PUT("/:product_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), productCont.UpdateProduct)
	productRouter.DELETE("/:product_id", productCont.DeleteProduct)

	//test

	// productRouter.POST("/", productCont.CreateProduct)
	// productRouter.GET("/", productCont.GetProducts)
	// productRouter.GET("/:product_id", productCont.GetProductByID)

	// productRouter.PUT("/:product_id", productCont.UpdateProduct)

}
