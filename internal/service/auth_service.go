package service

import (
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/internal/utils"
)

type IAuthService interface {
	SellerLogin(req *dto.LoginRequest) (*dto.Seller, string, string, error)
}

type AuthService struct {
	sellerRepository repository.ISellerRepository
}

func NewAuthService(sellerRepo repository.ISellerRepository) IAuthService {
	return AuthService{
		sellerRepository: sellerRepo,
	}
}

func (s AuthService) SellerLogin(req *dto.LoginRequest) (*dto.Seller, string, string, error) {
	sellerDTO, err := s.sellerRepository.GetSellerByUsername(req)
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

	return sellerDTO, accessToken, refreshToken, nil
}
