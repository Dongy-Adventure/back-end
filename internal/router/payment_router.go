package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r Router) AddPaymentRouter(rg *gin.RouterGroup) {

	cont := r.deps.PaymentController

	authRouter := rg.Group("payment")

	authRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), cont.HandlePayment)
}
