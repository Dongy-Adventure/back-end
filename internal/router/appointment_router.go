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

func (r Router) AddAppointmentRouter(rg *gin.RouterGroup, mongoDB *mongo.Database) {
	repo := repository.NewAppointmentRepository(mongoDB, "appointments")
	serv := service.NewAppointmentService(repo)
	cont := controller.NewAppointmentController(serv)

	appointmentRouter := rg.Group("appointment")

	appointmentRouter.POST("/", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.CreateAppointment)
	appointmentRouter.GET("/", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.GetAppointments)
	appointmentRouter.GET("/:appointment_id", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.GetAppointmentByID)
	appointmentRouter.GET("/order/:order_id", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.GetAppointmentByOrderID)
	appointmentRouter.PUT("/:appointment_id/date", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.UpdateAppointmentDate)
	appointmentRouter.PUT("/:appointment_id/place", middleware.JWTAuthMiddleWare(tokenmode.TokenMode.ACCESS_TOKEN), cont.UpdateAppointmentPlace)

}