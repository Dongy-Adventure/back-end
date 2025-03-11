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

func (r Router) AddOrderRouter(rg *gin.RouterGroup, mongoDB *mongo.Database) {
	repo := repository.NewOrderRepository(mongoDB, "orders")
	serv := service.NewOrderService(repo)
	cont := controller.NewOrderController(serv)

	orderRouter := rg.Group("order")

	orderRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), cont.CreateOrder)
	orderRouter.GET("/:user_id/:user_type", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), cont.GetOrdersByUserID)
	orderRouter.DELETE("/:order_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), cont.DeleteOrderByOrderID)
	orderRouter.PUT("/:order_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), cont.UpdateOrderByOrderID)
	orderRouter.PATCH("/:order_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), cont.UpdateOrderStatusByOrderID)

}
