package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
)

type IBuyerController interface {
	CreateBuyer(c *gin.Context)
	GetBuyerByID(c *gin.Context)
	UpdateBuyer(c *gin.Context)
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

func (s BuyerController) GetBuyerByID(c *gin.Context) {
	buyerID := c.Param("buyer_id")
	buyerDTO, err := s.buyerService.GetBuyerByID(buyerID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "No buyer with this buyerID",
			"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "get buyer success",
		"data":    buyerDTO,
	})
}

func (s BuyerController) UpdateBuyer(c *gin.Context) {
	buyerID := c.Param("id")

	var updatedBuyer model.Buyer
	if err := c.BindJSON(&updatedBuyer); err != nil {
	    c.JSON(http.StatusBadRequest, gin.H{
		   "error":   "Invalid request body, failed to bind JSON",
		   "message": err.Error(),
	    })
	    return
	}
 
	res, err := s.buyerService.UpdateBuyerData(buyerID, &updatedBuyer)
	if err != nil {
	    c.JSON(http.StatusInternalServerError, gin.H{
		   "error":   "Failed to update buyer data",
		   "message": err.Error(),
	    })
	    return
	}
 
	c.JSON(http.StatusOK, gin.H{
	    "message": "Update buyer success",
	    "data":    res,
	})
 }
 
