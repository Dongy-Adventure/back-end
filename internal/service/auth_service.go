package service

import (
	"fmt"

	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/internal/utils"
	"github.com/Dongy-s-Advanture/back-end/internal/utils/converter"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SellerLogin(req *dto.LoginRequest) (*dto.Seller, string, string, error)
	BuyerLogin(req *dto.LoginRequest) (*dto.Buyer, string, string, error)
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

	accessToken, accessTokenErr := utils.GenerateToken(tokenmode.TokenMode.ACCESS_TOKEN)
	if accessTokenErr != nil {
		return nil, "", "", accessTokenErr
	}
	refreshToken, refreshTokenErr := utils.GenerateToken(tokenmode.TokenMode.REFRESH_TOKEN)
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
	accessToken, accessTokenErr := utils.GenerateToken(tokenmode.TokenMode.ACCESS_TOKEN)
	if accessTokenErr != nil {
		return nil, "", "", accessTokenErr
	}
	refreshToken, refreshTokenErr := utils.GenerateToken(tokenmode.TokenMode.REFRESH_TOKEN)
	if refreshTokenErr != nil {
		return nil, "", "", refreshTokenErr
	}

	buyerDTO, err := converter.BuyerModelToDTO(buyerModel)
	if err != nil {
		return nil, "", "", err
	}

	return buyerDTO, accessToken, refreshToken, nil
}
