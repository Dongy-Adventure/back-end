package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	SellerLogin(*gin.Context)
}

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(s service.IAuthService) IAuthController {
	return AuthController{
		authService: s,
	}
}

func (a AuthController) SellerLogin(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body, failed to bind JSON",
			"message": err.Error(),
		})
		return
	}

	sellerDTO, accessToken, refreshToken, err := a.authService.SellerLogin(&req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Username or Password is incorrect",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":      "login success",
		"data":         sellerDTO,
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
