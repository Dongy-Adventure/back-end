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

func (r Router) AddReviewRouter(rg *gin.RouterGroup, mongoDB *mongo.Database) {
	sellerRepo := repository.NewSellerRepository(mongoDB, "sellers", "reviews")
	repo := repository.NewReviewRepository(mongoDB, "reviews", sellerRepo)
	serv := service.NewReviewService(repo)
	cont := controller.NewReviewController(serv)

	reviewRouter := rg.Group("review")

	reviewRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.CreateReview)
	reviewRouter.GET("/", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.GetReviews)
	reviewRouter.GET("/:review_id", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.GetReviewByID)
	reviewRouter.GET("/seller/:seller_id", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.GetReviewsBySellerID)
	reviewRouter.PUT("/:review_id", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.UpdateReview)

}