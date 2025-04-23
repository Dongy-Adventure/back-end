package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r Router) AddAuthRouter(rg *gin.RouterGroup) {

	cont := r.deps.AuthController

	authRouter := rg.Group("auth")

	authRouter.POST("/seller", cont.SellerLogin)
	authRouter.POST("/buyer", cont.BuyerLogin)
	authRouter.POST("/refresh", middleware.JWTAuthMiddleWare(tokenmode.REFRESH_TOKEN, r.deps.redis, r.deps.conf), cont.RefreshToken)
	authRouter.POST("/logout", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), cont.Logout)
}
