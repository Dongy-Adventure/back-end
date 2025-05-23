package controller

import (
	"net/http"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/service/auth"
	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Logout(c *gin.Context)
	SellerLogin(c *gin.Context)
	BuyerLogin(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type AuthController struct {
	authService auth.IAuthService
	config      *config.Config
}

func NewAuthController(c *config.Config, s auth.IAuthService) IAuthController {
	return AuthController{
		authService: s,
		config:      c,
	}
}

// SellerLogin godoc
//
//	@Summary		Seller login
//	@Description	Authenticate a seller and returns tokens
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			loginRequest	body		dto.LoginRequest	true	"Seller login credential"
//	@Success		200				{object}	dto.LoginResponse{data=dto.Seller}
//	@Failure		400				{object}	dto.ErrorResponse
//	@Failure		401				{object}	dto.ErrorResponse
//	@Failure		500				{object}	dto.ErrorResponse
//	@Router			/auth/seller/ [post]
func (a AuthController) SellerLogin(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	sellerDTO, accessToken, refreshToken, err := a.authService.SellerLogin(&req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusUnauthorized,
			Error:   "Username or Password is incorrect",
			Message: err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Success:               true,
		Status:                http.StatusOK,
		Message:               "login success",
		Data:                  sellerDTO,
		AccessToken:           accessToken,
		AccessTokenExpiredIn:  a.config.Auth.AccessTokenLifespanMinutes,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredIn: a.config.Auth.RefreshTokenLifespanMinutes,
	})
}

// BuyerLogin godoc
//
//	@Summary		Buyer login
//	@Description	Authenticate a buyer and returns tokens
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			loginRequest	body		dto.LoginRequest	true	"Buyer login credential"
//	@Success		200				{object}	dto.LoginResponse{data=dto.Buyer}
//	@Failure		400				{object}	dto.ErrorResponse
//	@Failure		401				{object}	dto.ErrorResponse
//	@Failure		500				{object}	dto.ErrorResponse
//	@Router			/auth/buyer/ [post]
func (a AuthController) BuyerLogin(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}

	buyerDTO, accessToken, refreshToken, err := a.authService.BuyerLogin(&req)

	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusUnauthorized,
			Error:   "Username or Password is incorrect",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Success:               true,
		Status:                http.StatusOK,
		Message:               "login success",
		Data:                  buyerDTO,
		AccessToken:           accessToken,
		AccessTokenExpiredIn:  a.config.Auth.AccessTokenLifespanMinutes,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredIn: a.config.Auth.RefreshTokenLifespanMinutes,
	})
}

// RefreshToken godoc
//
//	@Summary		Refresh token
//	@Description	Refresh access token for user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			refreshToken	header		string	true	"Bearer {refreshToken}"
//	@Success		200				{object}	dto.RefreshTokenResponse
//	@Failure		401				{object}	dto.ErrorResponse
//	@Router			/auth/refresh/ [post]
func (a AuthController) RefreshToken(c *gin.Context) {
	accessToken, err := a.authService.RefreshToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Success: false, Status: http.StatusUnauthorized, Message: err.Error(), Error: "Unauthorized"})
		return
	}
	c.JSON(http.StatusOK, dto.RefreshTokenResponse{Success: true, Status: http.StatusOK, Message: "Refresh success", AccessToken: accessToken, AccessTokenExpiredIn: a.config.Auth.AccessTokenLifespanMinutes})

}

// Logout godoc
//
//	@Summary		User logout
//	@Description	Invalidate user's token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			logoutRequest	body		dto.LogoutRequest	true	"User's tokens"
//	@Success		200				{object}	dto.SuccessResponse{data=string}
//	@Failure		400				{object}	dto.ErrorResponse
//	@Failure		500				{object}	dto.ErrorResponse
//	@Router			/auth/logout/ [post]
func (a AuthController) Logout(c *gin.Context) {
	var req dto.LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Error:   "Invalid request body, failed to bind JSON",
			Message: err.Error(),
		})
		return
	}
	err := a.authService.Logout(req.AccessToken, req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Error:   "Something is wrong",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dto.SuccessResponse{Success: true, Status: http.StatusOK, Message: "logout success", Data: "invalidate access and refresh token success"})

}
