package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/controller"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Router) AddProductRouter(rg *gin.RouterGroup, mongoDB *mongo.Database) {
	// Repositories
	productRepo := repository.NewProductRepository(mongoDB, "products")
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	// Product Router Group
	productRouter := rg.Group("products")

	// Define Routes
	productRouter.POST("/", productController.CreateProduct)
	productRouter.GET("/:id", productController.GetProductByID)
	productRouter.GET("/", productController.GetAllProducts)
	productRouter.PUT("/:id", productController.UpdateProduct)
	productRouter.DELETE("/:id", productController.DeleteProduct)
}
