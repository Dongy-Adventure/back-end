package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/repository"
	"github.com/Dongy-s-Advanture/back-end/pkg/redis"
	"github.com/Dongy-s-Advanture/back-end/pkg/utils/converter"
	"github.com/Dongy-s-Advanture/back-end/pkg/utils/token"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	SellerLogin(req *dto.LoginRequest) (*dto.Seller, string, string, error)
	BuyerLogin(req *dto.LoginRequest) (*dto.Buyer, string, string, error)
	RefreshToken(c *gin.Context) (string, error)
	InvalidateToken(token string, expirationTime time.Duration) error
	Logout(accessToken string, refreshToken string) error
}

type AuthService struct {
	conf             *config.Config
	redisDB          redis.IRedisClient
	sellerRepository repository.ISellerRepository
	buyerRepository  repository.IBuyerRepository
}

func NewAuthService(conf *config.Config, redisDB redis.IRedisClient, sellerRepo repository.ISellerRepository, buyerRepo repository.IBuyerRepository) IAuthService {
	return AuthService{
		conf:             conf,
		redisDB:          redisDB,
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

	accessToken, accessTokenErr := token.GenerateToken(s.conf, sellerModel.SellerID.Hex(), tokenmode.ACCESS_TOKEN)
	if accessTokenErr != nil {
		return nil, "", "", accessTokenErr
	}
	refreshToken, refreshTokenErr := token.GenerateToken(s.conf, sellerModel.SellerID.Hex(), tokenmode.REFRESH_TOKEN)
	if refreshTokenErr != nil {
		return nil, "", "", refreshTokenErr
	}

	sellerDTO, _ := converter.SellerModelToDTO(sellerModel)

	return sellerDTO, accessToken, refreshToken, nil
}

func (s AuthService) BuyerLogin(req *dto.LoginRequest) (*dto.Buyer, string, string, error) {
	buyerModel, err := s.buyerRepository.GetBuyerByUsername(req)
	if err != nil {
		return nil, "", "", err
	}
	accessToken, accessTokenErr := token.GenerateToken(s.conf, buyerModel.BuyerID.Hex(), tokenmode.ACCESS_TOKEN)
	if accessTokenErr != nil {
		return nil, "", "", accessTokenErr
	}
	refreshToken, refreshTokenErr := token.GenerateToken(s.conf, buyerModel.BuyerID.Hex(), tokenmode.REFRESH_TOKEN)
	if refreshTokenErr != nil {
		return nil, "", "", refreshTokenErr
	}

	buyerDTO, _ := converter.BuyerModelToDTO(buyerModel)

	return buyerDTO, accessToken, refreshToken, nil
}

func (s AuthService) RefreshToken(c *gin.Context) (string, error) {
	tkn, err := token.ValidateToken(c, s.conf, s.redisDB, tokenmode.REFRESH_TOKEN)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}
	userID, err := token.ExtractID(tkn)
	if err != nil {
		return "", fmt.Errorf("no userID in token")
	}
	accessToken, _ := token.GenerateToken(s.conf, userID, tokenmode.ACCESS_TOKEN)

	return accessToken, nil
}

func (s AuthService) InvalidateToken(token string, expirationTime time.Duration) error {
	ctx := context.Background()
	err := s.redisDB.SetEx(ctx, "blacklist:"+token, "invalid", expirationTime).Err()
	if err != nil {
		return fmt.Errorf("could not invalidate token: %v", err)
	}
	return nil
}

func (s AuthService) Logout(accessToken string, refreshToken string) error {
	accessTokenExpiredIn := s.conf.Auth.AccessTokenLifespanMinutes
	refreshTokenExpiredIn := s.conf.Auth.RefreshTokenLifespanMinutes

	if err := s.InvalidateToken(accessToken, time.Minute*time.Duration(accessTokenExpiredIn)); err != nil {
		return err
	}
	if err := s.InvalidateToken(refreshToken, time.Minute*time.Duration(refreshTokenExpiredIn)); err != nil {
		return err
	}
	return nil
}
