package router

import (
	"log"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/Dongy-s-Advanture/back-end/internal/controller"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/Dongy-s-Advanture/back-end/pkg/utils"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/omise/omise-go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dependencies struct {
	BuyerRepo       repository.IBuyerRepository
	BuyerService    service.IBuyerService
	BuyerController controller.IBuyerController

	SellerRepo       repository.ISellerRepository
	SellerService    service.ISellerService
	SellerController controller.ISellerController

	AuthService    service.IAuthService
	AuthController controller.IAuthController

	ProductRepo       repository.IProductRepository
	ProductService    service.IProductService
	ProductController controller.IProductController

	ReviewRepo       repository.IReviewRepository
	ReviewService    service.IReviewService
	ReviewController controller.IReviewController

	AppointmentRepo       repository.IAppointmentRepository
	AppointmentService    service.IAppointmentService
	AppointmentController controller.IAppointmentController

	OrderRepo       repository.IOrderRepository
	OrderService    service.IOrderService
	OrderController controller.IOrderController

	PaymentService    service.IPaymentService
	PaymentController controller.IPaymentController

	AdvertisementRepo       repository.IAdvertisementRepository
	AdvertisementService    service.IAdvertisementService
	AdvertisementController controller.IAdvertisementController
}

func NewDependencies(mongoDB *mongo.Database, redisDB *redis.Client, s3Client *s3.Client, conf *config.Config) *Dependencies {

	// Initialize third party
	omiseClient, e := omise.NewClient(conf.Payment.Public, conf.Payment.Private)
	if e != nil {
		log.Fatal(e)
	}

	s3Service := utils.NewS3Service(s3Client, &conf.AWS)

	// Initialize repositories
	buyerRepo := repository.NewBuyerRepository(mongoDB, "buyers")
	sellerRepo := repository.NewSellerRepository(mongoDB, "sellers", "reviews")
	productRepo := repository.NewProductRepository(mongoDB, *s3Service, "products")
	reviewRepo := repository.NewReviewRepository(mongoDB, "reviews", sellerRepo)
	appointmentRepo := repository.NewAppointmentRepository(mongoDB, "appointments")
	orderRepo := repository.NewOrderRepository(mongoDB, "orders")
	advertisementRepo := repository.NewAdvertisementRepository(mongoDB, *s3Service, "advertisements")

	// Initialize services
	buyerService := service.NewBuyerService(buyerRepo)
	sellerService := service.NewSellerService(sellerRepo)
	authService := service.NewAuthService(conf, redisDB, sellerRepo, buyerRepo)
	productService := service.NewProductService(productRepo)
	reviewService := service.NewReviewService(reviewRepo)
	appointmentService := service.NewAppointmentService(appointmentRepo)
	orderService := service.NewOrderService(orderRepo, appointmentRepo, sellerRepo, productRepo)
	paymentService := service.NewPaymentService(omiseClient)
	advertisementService := service.NewAdvertisementService(advertisementRepo)

	// Initialize controllers
	buyerController := controller.NewBuyerController(buyerService)
	sellerController := controller.NewSellerController(sellerService)
	authController := controller.NewAuthController(conf, authService)
	productController := controller.NewProductController(productService)
	reviewController := controller.NewReviewController(reviewService)
	appointmentController := controller.NewAppointmentController(appointmentService)
	orderController := controller.NewOrderController(orderService, paymentService)
	paymentController := controller.NewPaymentController(paymentService)
	advertisementController := controller.NewAdvertisementController(advertisementService)

	return &Dependencies{
		BuyerRepo:       buyerRepo,
		BuyerService:    buyerService,
		BuyerController: buyerController,

		SellerRepo:       sellerRepo,
		SellerService:    sellerService,
		SellerController: sellerController,

		AuthService:    authService,
		AuthController: authController,

		ProductRepo:       productRepo,
		ProductService:    productService,
		ProductController: productController,

		ReviewRepo:       reviewRepo,
		ReviewService:    reviewService,
		ReviewController: reviewController,

		AppointmentRepo:       appointmentRepo,
		AppointmentService:    appointmentService,
		AppointmentController: appointmentController,

		OrderRepo:       orderRepo,
		OrderService:    orderService,
		OrderController: orderController,

		PaymentService:    paymentService,
		PaymentController: paymentController,

		AdvertisementRepo: advertisementRepo,
		AdvertisementService: advertisementService,
		AdvertisementController: advertisementController,

	}
}
