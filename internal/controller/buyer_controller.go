package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
)

type IBuyerController interface {
	CreateBuyer(c *gin.Context)
}

type BuyerController struct {
	buyerService service.IBuyerService
}

func NewBuyerController(s service.IBuyerService) IBuyerController {
	return BuyerController{
		buyerService: s,
	}
}

func (s BuyerController) CreateBuyer(c *gin.Context) {
	var newBuyer model.Buyer

	if err := c.BindJSON(&newBuyer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body, failed to bind JSON",
			"message": err.Error(),
		})
		return
	}
	res, err := s.buyerService.CreateBuyerData(&newBuyer)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to insert to database",
			"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created buyer success",
		"data":    res,
	})
}
