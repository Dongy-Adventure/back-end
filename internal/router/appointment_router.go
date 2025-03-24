package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Router) AddAppointmentRouter(rg *gin.RouterGroup, mongoDB *mongo.Database) {

	appointmentRouter := rg.Group("appointment")

	appointmentRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), appointmentCont.CreateAppointment)
	appointmentRouter.GET("/", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), appointmentCont.GetAppointments)
	appointmentRouter.GET("/:appointment_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), appointmentCont.GetAppointmentByID)
	appointmentRouter.GET("/order/:order_id", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), appointmentCont.GetAppointmentByOrderID)
	appointmentRouter.PUT("/:appointment_id/date", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), appointmentCont.UpdateAppointmentDate)
	appointmentRouter.PUT("/:appointment_id/place", middleware.JWTAuthMiddleWare(tokenmode.ACCESS_TOKEN), appointmentCont.UpdateAppointmentPlace)

}
