package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r Router) AddAppointmentRouter(rg *gin.RouterGroup) {

	appointmentCont := r.deps.AppointmentController
	appointmentRouter := rg.Group("appointment")

	appointmentRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), appointmentCont.CreateAppointment)
	appointmentRouter.GET("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), appointmentCont.GetAppointments)
	appointmentRouter.GET("/:appointment_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), appointmentCont.GetAppointmentByID)
	appointmentRouter.GET("/order/:order_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), appointmentCont.GetAppointmentByOrderID)
	appointmentRouter.PUT("/:appointment_id/date", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), appointmentCont.UpdateAppointmentDate)
	appointmentRouter.PUT("/:appointment_id/place", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN, r.deps.redis, r.deps.conf), appointmentCont.UpdateAppointmentPlace)

}
