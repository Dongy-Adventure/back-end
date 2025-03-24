package router

import (
	"fmt"

	"time"

	docs "github.com/Dongy-s-Advanture/back-end/docs"
	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/Dongy-s-Advanture/back-end/internal/controller"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.mongodb.org/mongo-driver/mongo"
)

var buyerRepo repository.IBuyerRepository
var buyerServ service.IBuyerService
var buyerCont controller.IBuyerController

var sellerRepo repository.ISellerRepository
var sellerServ service.ISellerService
var sellerCont controller.ISellerController

var productRepo repository.IProductRepository
var productServ service.IProductService
var productCont controller.IProductController

var reviewRepo repository.IReviewRepository
var reviewServ service.IReviewService
var reviewCont controller.IReviewController

var appointmentRepo repository.IAppointmentRepository
var appointmentServ service.IAppointmentService
var appointmentCont controller.IAppointmentController

var orderRepo repository.IOrderRepository
var orderServ service.IOrderService
var orderCont controller.IOrderController

type Router struct {
	g    *gin.Engine
	conf *config.Config
}

func NewRouter(g *gin.Engine, conf *config.Config) *Router {
	return &Router{g, conf}
}

func setUp(mongoDB *mongo.Database, redisDB *redis.Client) {

	buyerRepo = repository.NewBuyerRepository(mongoDB, "buyers")
	buyerServ = service.NewBuyerService(buyerRepo)
	buyerCont = controller.NewBuyerController(buyerServ)

	sellerRepo = repository.NewSellerRepository(mongoDB, "sellers", "reviews")
	sellerServ = service.NewSellerService(sellerRepo)
	sellerCont = controller.NewSellerController(sellerServ)

	productRepo = repository.NewProductRepository(mongoDB, "products")
	productServ = service.NewProductService(productRepo)
	productCont = controller.NewProductController(productServ)

	reviewRepo = repository.NewReviewRepository(mongoDB, "reviews", sellerRepo)
	reviewServ = service.NewReviewService(reviewRepo)
	reviewCont = controller.NewReviewController(reviewServ)

	appointmentRepo = repository.NewAppointmentRepository(mongoDB, "appointments")
	appointmentServ = service.NewAppointmentService(appointmentRepo)
	appointmentCont = controller.NewAppointmentController(appointmentServ)

	orderRepo = repository.NewOrderRepository(mongoDB, "orders")
	orderServ = service.NewOrderService(orderRepo, appointmentRepo, sellerRepo)
	orderCont = controller.NewOrderController(orderServ)
}

func (r *Router) Run(mongoDB *mongo.Database, redisDB *redis.Client) {

	// CORS setting
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowMethods = []string{"OPTIONS", "PATCH", "PUT", "GET", "POST", "DELETE"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"} // Allow Authorization header
	corsConfig.AllowCredentials = true                                  // If you are using cookies or Authorization header

	// Optional: Handle preflight cache
	corsConfig.MaxAge = 12 * time.Hour

	r.g.Use(cors.New(corsConfig))

	r.g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Swagger setting
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// versioning
	v1 := r.g.Group("/api/v1")

	// setup
	setUp(mongoDB, redisDB)

	// Add related path
	r.AddSellerRouter(v1, mongoDB)
	r.AddBuyerRouter(v1, mongoDB)
	r.AddAuthRouter(v1, mongoDB, redisDB)
	r.AddProductRouter(v1, mongoDB)
	r.AddOrderRouter(v1, mongoDB)
	r.AddReviewRouter(v1, mongoDB)
	r.AddAppointmentRouter(v1, mongoDB)

	err := r.g.Run(":" + r.conf.App.Port)
	if err != nil {
		panic(fmt.Sprintf("Failed to run the server : %v", err))
	}
}
