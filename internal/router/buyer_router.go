package router

import (
	"github.com/Dongy-s-Advanture/back-end/internal/controller"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r Router) AddBuyerRouter(rg *gin.RouterGroup, mongoDB *mongo.Database) {
	repo := repository.NewBuyerRepository(mongoDB, "buyers")
	serv := service.NewBuyerService(repo)
	cont := controller.NewBuyerController(serv)

	buyerRouter := rg.Group("buyer")

	buyerRouter.POST("/", cont.CreateBuyer)
	buyerRouter.GET("/:buyer_id",cont.GetBuyerByID)
	buyerRouter.PUT("/:buyer_id",cont.UpdateBuyer)
}
