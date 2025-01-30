package service

import (
	"fmt"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/internal/utils/converter"
	"github.com/Dongy-s-Advanture/back-end/internal/utils/token"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SellerLogin(req *dto.LoginRequest) (*dto.Seller, string, string, error)
	BuyerLogin(req *dto.LoginRequest) (*dto.Buyer, string, string, error)
	RefreshToken(c *gin.Context) (string, error)
}

type AuthService struct {
	sellerRepository repository.ISellerRepository
	buyerRepository  repository.IBuyerRepository
}

func NewAuthService(sellerRepo repository.ISellerRepository, buyerRepo repository.IBuyerRepository) IAuthService {
	return AuthService{
		sellerRepository: sellerRepo,
		buyerRepository:  buyerRepo,
	}
}

func (s AuthService) SellerLogin(req *dto.LoginRequest) (*dto.Seller, string, string, error) {
	sellerModel, err := s.sellerRepository.GetSellerByUsername(req)
	if err != nil {
		return nil, "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(sellerModel.Password), []byte(req.Password))
	if err != nil {
		return nil, "", "", fmt.Errorf("invalid username or password")
	}

	accessToken, accessTokenErr := token.GenerateToken(tokenmode.TokenMode.ACCESS_TOKEN)
	if accessTokenErr != nil {
		return nil, "", "", accessTokenErr
	}
	refreshToken, refreshTokenErr := token.GenerateToken(tokenmode.TokenMode.REFRESH_TOKEN)
	if refreshTokenErr != nil {
		return nil, "", "", refreshTokenErr
	}

	sellerDTO, err := converter.SellerModelToDTO(sellerModel)
	if err != nil {
		return nil, "", "", err
	}

	return sellerDTO, accessToken, refreshToken, nil
}

func (s AuthService) BuyerLogin(req *dto.LoginRequest) (*dto.Buyer, string, string, error) {
	buyerModel, err := s.buyerRepository.GetBuyerByUsername(req)
	if err != nil {
		return nil, "", "", err
	}
	accessToken, accessTokenErr := token.GenerateToken(tokenmode.TokenMode.ACCESS_TOKEN)
	if accessTokenErr != nil {
		return nil, "", "", accessTokenErr
	}
	refreshToken, refreshTokenErr := token.GenerateToken(tokenmode.TokenMode.REFRESH_TOKEN)
	if refreshTokenErr != nil {
		return nil, "", "", refreshTokenErr
	}

	buyerDTO, err := converter.BuyerModelToDTO(buyerModel)
	if err != nil {
		return nil, "", "", err
	}

	return buyerDTO, accessToken, refreshToken, nil
}

func (s AuthService) RefreshToken(c *gin.Context) (string, error) {
	err := token.ValidateToken(c, tokenmode.TokenMode.REFRESH_TOKEN)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}
	accessToken, accessTokenErr := token.GenerateToken(tokenmode.TokenMode.ACCESS_TOKEN)
	if accessTokenErr != nil {
		return "", accessTokenErr
	}
	return accessToken, nil
}
