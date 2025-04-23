package router

import (
	"github.com/gin-gonic/gin"
)

func (r Router) AddPaymentRouter(rg *gin.RouterGroup) {

	cont := r.deps.PaymentController

	paymentRouter := rg.Group("payment")

	paymentRouter.POST("/", cont.HandlePayment)
	// paymentRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), cont.HandlePayment)
	paymentRouter.GET("/sse/:charge_id", cont.SSEHandler)
	paymentRouter.POST("/webhooks/omise", cont.OmiseWebhookHandler)

}
