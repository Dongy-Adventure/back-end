package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
)

type ISellerController interface {
	CreateSeller(c *gin.Context)
}

type SellerController struct {
	sellerService service.ISellerService
}

func NewSellerController(s service.ISellerService) ISellerController {
	return SellerController{
		sellerService: s,
	}
}

func (s SellerController) CreateSeller(c *gin.Context) {
	var newSeller model.Seller

	if err := c.BindJSON(&newSeller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body, failed to bind JSON",
			"message": err.Error(),
		})
		return
	}
	res, err := s.sellerService.CreateSellerData(&newSeller)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to insert to database",
			"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created seller success",
		"data":    res,
	})
}
