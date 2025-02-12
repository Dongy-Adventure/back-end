package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/controller"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Router) AddAuthRouter(rg *gin.RouterGroup, mongoDB *mongo.Database, redisDB *redis.Client) {
	sellerRepo := repository.NewSellerRepository(mongoDB, "sellers", "reviews")
	buyerRepo := repository.NewBuyerRepository(mongoDB, "buyers")

	serv := service.NewAuthService(r.conf, redisDB, sellerRepo, buyerRepo)
	cont := controller.NewAuthController(serv, r.conf)

	authRouter := rg.Group("auth")

	authRouter.POST("/seller", cont.SellerLogin)
	authRouter.POST("/buyer", cont.BuyerLogin)
	authRouter.POST("/refresh", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.REFRESH_TOKEN), cont.RefreshToken)
	authRouter.POST("/logout", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.Logout)
}
