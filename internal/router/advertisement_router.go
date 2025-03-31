package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r Router) AddAdvertisementRouter(rg *gin.RouterGroup) {

	advertisementCont := r.deps.AdvertisementController
	advertisementRouter := rg.Group("advertisement")

	advertisementRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), advertisementCont.CreateAdvertisement)
	advertisementRouter.GET("/", advertisementCont.GetAdvertisements)
	advertisementRouter.GET("/:advertisement_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), advertisementCont.GetAdvertisementByID)
	advertisementRouter.GET("/seller/:seller_id", advertisementCont.GetAdvertisementsBySellerID)
	advertisementRouter.GET("/product/:product_id", advertisementCont.GetAdvertisementsByProductID)
	advertisementRouter.PUT("/:advertisement_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), advertisementCont.UpdateAdvertisement)
	advertisementRouter.DELETE("/:advertisement_id", advertisementCont.DeleteAdvertisement)

}
