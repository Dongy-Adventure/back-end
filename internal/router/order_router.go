package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r Router) AddOrderRouter(rg *gin.RouterGroup) {

	orderCont := r.deps.OrderController
	orderRouter := rg.Group("order")

	orderRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis), orderCont.CreateOrder)
	orderRouter.GET("/:user_id/:user_type", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis), orderCont.GetOrdersByUserID)
	orderRouter.DELETE("/:order_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis), orderCont.DeleteOrderByOrderID)
	orderRouter.PUT("/:order_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis), orderCont.UpdateOrderByOrderID)
	orderRouter.PATCH("/:order_id", orderCont.UpdateOrderStatusByOrderID)

}
