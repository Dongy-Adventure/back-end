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

func (r Router) AddProductRouter(rg *gin.RouterGroup, mongoDB *mongo.Database) {
	repo := repository.NewProductRepository(mongoDB, "products")
	serv := service.NewProductService(repo)
	cont := controller.NewProductController(serv)

	productRouter := rg.Group("product")

	productRouter.POST("/", cont.CreateProduct)
	productRouter.GET("/", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.GetProducts)
	productRouter.GET("/:product_id", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.GetProductByID)
	productRouter.PUT("/:product_id", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.UpdateProduct)
}
