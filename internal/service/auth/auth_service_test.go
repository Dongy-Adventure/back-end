package auth

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	mocks "github.com/Dongy-s-Advanture/back-end/pkg/mock/repository"
	"github.com/Dongy-s-Advanture/back-end/pkg/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthService_SellerLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSellerRepo := mocks.NewMockISellerRepository(ctrl)
	mockBuyerRepo := mocks.NewMockIBuyerRepository(ctrl)
	mockRedis := mocks.NewMockIRedisClient(ctrl)

	conf := &config.Config{
		Auth: config.AuthConfig{
			AccessTokenLifespanMinutes:  15,
			RefreshTokenLifespanMinutes: 1440,
			AccessTokenSecret:           "test-secret",
			RefreshTokenSecret:          "test-secret",
		},
	}

	authService := NewAuthService(conf, mockRedis, mockSellerRepo, mockBuyerRepo)

	t.Run("successful seller login", func(t *testing.T) {
		req := &dto.LoginRequest{
			Username: "test-seller",
			Password: "password123",
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		sellerModel := &model.Seller{
			SellerID: primitive.NewObjectID(),
			Username: req.Username,
			Password: string(hashedPassword),
		}

		mockSellerRepo.EXPECT().GetSellerByUsername(req).Return(sellerModel, nil)

		sellerDTO, accessToken, refreshToken, err := authService.SellerLogin(req)
		assert.NoError(t, err)
		assert.NotEmpty(t, accessToken)
		assert.NotEmpty(t, refreshToken)
		assert.Equal(t, sellerModel.Username, sellerDTO.Username)
	})

	t.Run("invalid username or password", func(t *testing.T) {
		req := &dto.LoginRequest{
			Username: "test-seller",
			Password: "wrong-password",
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correct-password"), bcrypt.DefaultCost)
		sellerModel := &model.Seller{
			SellerID: primitive.NewObjectID(),
			Username: req.Username,
			Password: string(hashedPassword),
		}

		mockSellerRepo.EXPECT().GetSellerByUsername(req).Return(sellerModel, nil)

		_, _, _, err := authService.SellerLogin(req)
		assert.Error(t, err)
		assert.Equal(t, "invalid username or password", err.Error())
	})

	t.Run("seller not found", func(t *testing.T) {
		req := &dto.LoginRequest{
			Username: "nonexistent-seller",
			Password: "password123",
		}

		mockSellerRepo.EXPECT().GetSellerByUsername(req).Return(nil, errors.New("seller not found"))

		_, _, _, err := authService.SellerLogin(req)
		assert.Error(t, err)
		assert.Equal(t, "seller not found", err.Error())
	})

	t.Run("access token lifespan is not set", func(t *testing.T) {
		// Config with zero access token lifespan
		invalidConf := &config.Config{
			Auth: config.AuthConfig{
				AccessTokenLifespanMinutes:  0, // This should trigger the error
				RefreshTokenLifespanMinutes: 1440,
				AccessTokenSecret:           "test-secret",
			},
		}

		// Create service with invalid config
		authService := NewAuthService(invalidConf, mockRedis, mockSellerRepo, mockBuyerRepo)

		req := &dto.LoginRequest{
			Username: "test-seller",
			Password: "password123",
		}

		// Mock a successful seller lookup first
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		sellerModel := &model.Seller{
			SellerID: primitive.NewObjectID(),
			Username: req.Username,
			Password: string(hashedPassword),
		}
		mockSellerRepo.EXPECT().GetSellerByUsername(req).Return(sellerModel, nil)

		// Execute
		_, _, _, err := authService.SellerLogin(req)

		// Verify
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error generating access token")
	})

	t.Run("refresh token generation error", func(t *testing.T) {

		invalidConf := &config.Config{
			Auth: config.AuthConfig{
				AccessTokenLifespanMinutes:  15,
				RefreshTokenLifespanMinutes: 0,
				AccessTokenSecret:           "test-secret",
			},
		}
		authService := NewAuthService(invalidConf, mockRedis, mockSellerRepo, mockBuyerRepo)

		req := &dto.LoginRequest{
			Username: "test-seller",
			Password: "password123",
		}

		// Mock a successful seller lookup first
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		sellerModel := &model.Seller{
			SellerID: primitive.NewObjectID(),
			Username: req.Username,
			Password: string(hashedPassword),
		}
		mockSellerRepo.EXPECT().GetSellerByUsername(req).Return(sellerModel, nil)

		// Execute
		_, _, _, err := authService.SellerLogin(req)

		// Verify
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error generating refresh token")
	})
}

func TestAuthService_BuyerLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSellerRepo := mocks.NewMockISellerRepository(ctrl)
	mockBuyerRepo := mocks.NewMockIBuyerRepository(ctrl)
	mockRedis := mocks.NewMockIRedisClient(ctrl)

	conf := &config.Config{
		Auth: config.AuthConfig{
			AccessTokenLifespanMinutes:  15,
			RefreshTokenLifespanMinutes: 1440,
			AccessTokenSecret:           "test-secret",
			RefreshTokenSecret:          "test-secret",
		},
	}

	authService := NewAuthService(conf, mockRedis, mockSellerRepo, mockBuyerRepo)

	t.Run("successful buyer login", func(t *testing.T) {
		req := &dto.LoginRequest{
			Username: "test-buyer",
			Password: "password123",
		}

		buyerModel := &model.Buyer{
			BuyerID:  primitive.NewObjectID(),
			Username: req.Username,
			Password: "hashed-password", // Password validation is not implemented in buyer login
		}

		mockBuyerRepo.EXPECT().GetBuyerByUsername(req).Return(buyerModel, nil)

		buyerDTO, accessToken, refreshToken, err := authService.BuyerLogin(req)
		assert.NoError(t, err)
		assert.NotEmpty(t, accessToken)
		assert.NotEmpty(t, refreshToken)
		assert.Equal(t, buyerModel.Username, buyerDTO.Username)
	})

	t.Run("buyer not found", func(t *testing.T) {
		req := &dto.LoginRequest{
			Username: "nonexistent-buyer",
			Password: "password123",
		}

		mockBuyerRepo.EXPECT().GetBuyerByUsername(req).Return(nil, errors.New("buyer not found"))

		_, _, _, err := authService.BuyerLogin(req)
		assert.Error(t, err)
		assert.Equal(t, "buyer not found", err.Error())
	})

	t.Run("access token lifespan is not set", func(t *testing.T) {
		// Config with zero access token lifespan
		invalidConf := &config.Config{
			Auth: config.AuthConfig{
				AccessTokenLifespanMinutes:  0, // This should trigger the error
				RefreshTokenLifespanMinutes: 1440,
				AccessTokenSecret:           "test-secret",
			},
		}

		// Create service with invalid config
		authService := NewAuthService(invalidConf, mockRedis, mockSellerRepo, mockBuyerRepo)

		req := &dto.LoginRequest{
			Username: "test-seller",
			Password: "password123",
		}

		// Mock a successful seller lookup first
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		buyerModel := &model.Buyer{
			BuyerID:  primitive.NewObjectID(),
			Username: req.Username,
			Password: string(hashedPassword),
		}
		mockBuyerRepo.EXPECT().GetBuyerByUsername(req).Return(buyerModel, nil)

		// Execute
		_, _, _, err := authService.BuyerLogin(req)

		// Verify
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error generating access token")
	})

	t.Run("refresh token generation error", func(t *testing.T) {

		invalidConf := &config.Config{
			Auth: config.AuthConfig{
				AccessTokenLifespanMinutes:  15,
				RefreshTokenLifespanMinutes: 0,
				AccessTokenSecret:           "test-secret",
			},
		}
		authService := NewAuthService(invalidConf, mockRedis, mockSellerRepo, mockBuyerRepo)

		req := &dto.LoginRequest{
			Username: "test-seller",
			Password: "password123",
		}

		// Mock a successful seller lookup first
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		buyerModel := &model.Buyer{
			BuyerID:  primitive.NewObjectID(),
			Username: req.Username,
			Password: string(hashedPassword),
		}
		mockBuyerRepo.EXPECT().GetBuyerByUsername(req).Return(buyerModel, nil)

		// Execute
		_, _, _, err := authService.BuyerLogin(req)

		// Verify
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error generating refresh token")
	})
}

func TestAuthService_RefreshToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSellerRepo := mocks.NewMockISellerRepository(ctrl)
	mockBuyerRepo := mocks.NewMockIBuyerRepository(ctrl)
	mockRedis := mocks.NewMockIRedisClient(ctrl)

	conf := &config.Config{
		Auth: config.AuthConfig{
			AccessTokenLifespanMinutes:  15,
			RefreshTokenLifespanMinutes: 1440,
			AccessTokenSecret:           "test-secret",
			RefreshTokenSecret:          "test-secret",
		},
	}

	authService := NewAuthService(conf, mockRedis, mockSellerRepo, mockBuyerRepo)

	t.Run("successful token refresh", func(t *testing.T) {
		userID := primitive.NewObjectID()
		refreshToken, _ := token.GenerateToken(conf, userID.String(), tokenmode.REFRESH_TOKEN)
		mockRedis.
			EXPECT().
			Exists(gomock.Any(), gomock.Any()).
			Return(redis.NewIntCmd(context.Background())).
			Times(2)

		c := &gin.Context{}
		c.Request = &http.Request{
			Header: http.Header{
				"Authorization": []string{"Bearer " + refreshToken},
			},
		}

		newAccessToken, err := authService.RefreshToken(c)
		assert.NoError(t, err)
		assert.NotEmpty(t, newAccessToken)
	})

	t.Run("invalid refresh token", func(t *testing.T) {
		c := &gin.Context{}
		c.Request = &http.Request{
			Header: http.Header{
				"Authorization": []string{"Bearer invalid-token"},
			},
		}

		_, err := authService.RefreshToken(c)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid refresh token")
	})

	t.Run("missing userID in token", func(t *testing.T) {
		// Create a valid token WITHOUT userID
		claims := jwt.MapClaims{
			"exp": time.Now().Add(time.Minute * 15).Unix(),
		}
		tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenStr, _ := tokenObj.SignedString([]byte(conf.Auth.RefreshTokenSecret))

		c := &gin.Context{}
		c.Request = &http.Request{
			Header: http.Header{
				"Authorization": []string{"Bearer " + tokenStr},
			},
		}

		_, err := authService.RefreshToken(c)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no userID in token")
	})

}

func TestAuthService_Logout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSellerRepo := mocks.NewMockISellerRepository(ctrl)
	mockBuyerRepo := mocks.NewMockIBuyerRepository(ctrl)
	mockRedis := mocks.NewMockIRedisClient(ctrl)

	conf := &config.Config{
		Auth: config.AuthConfig{
			AccessTokenLifespanMinutes:  15,
			RefreshTokenLifespanMinutes: 1440,
			AccessTokenSecret:           "test-secret",
			RefreshTokenSecret:          "test-secret",
		},
	}

	authService := NewAuthService(conf, mockRedis, mockSellerRepo, mockBuyerRepo)

	t.Run("successful logout", func(t *testing.T) {
		accessToken := "test-access-token"
		refreshToken := "test-refresh-token"

		mockRedis.EXPECT().SetEx(context.Background(), "blacklist:"+accessToken, "invalid", time.Minute*15).Return(redis.NewStatusResult("OK", nil))
		mockRedis.EXPECT().SetEx(context.Background(), "blacklist:"+refreshToken, "invalid", time.Minute*1440).Return(redis.NewStatusResult("OK", nil))

		err := authService.Logout(accessToken, refreshToken)
		assert.NoError(t, err)
	})

	t.Run("failed to invalidate access token", func(t *testing.T) {
		accessToken := "test-access-token"
		refreshToken := "test-refresh-token"

		mockRedis.EXPECT().SetEx(context.Background(), "blacklist:"+accessToken, "invalid", time.Minute*15).Return(redis.NewStatusResult("", errors.New("redis error")))

		err := authService.Logout(accessToken, refreshToken)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "could not invalidate token")
	})

	t.Run("failed to invalidate refresh token", func(t *testing.T) {
		accessToken := "test-access-token"
		refreshToken := "test-refresh-token"

		mockRedis.EXPECT().SetEx(context.Background(), "blacklist:"+accessToken, "invalid", time.Minute*15).Return(redis.NewStatusResult("OK", nil))
		mockRedis.EXPECT().SetEx(context.Background(), "blacklist:"+refreshToken, "invalid", time.Minute*1440).Return(redis.NewStatusResult("", errors.New("redis error")))

		err := authService.Logout(accessToken, refreshToken)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "could not invalidate token")
	})
}

func TestAuthService_invalidateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSellerRepo := mocks.NewMockISellerRepository(ctrl)
	mockBuyerRepo := mocks.NewMockIBuyerRepository(ctrl)
	mockRedis := mocks.NewMockIRedisClient(ctrl)

	conf := &config.Config{
		Auth: config.AuthConfig{
			AccessTokenLifespanMinutes:  15,
			RefreshTokenLifespanMinutes: 1440,
			AccessTokenSecret:           "test-secret",
			RefreshTokenSecret:          "test-secret",
		},
	}

	authService := NewAuthService(conf, mockRedis, mockSellerRepo, mockBuyerRepo)

	t.Run("successful token invalidation", func(t *testing.T) {
		token := "test-token"
		expiration := time.Minute * 15

		mockRedis.EXPECT().SetEx(context.Background(), "blacklist:"+token, "invalid", expiration).Return(redis.NewStatusResult("OK", nil))

		err := authService.InvalidateToken(token, expiration)
		assert.NoError(t, err)
	})

	t.Run("failed token invalidation", func(t *testing.T) {
		token := "test-token"
		expiration := time.Minute * 15

		mockRedis.EXPECT().SetEx(context.Background(), "blacklist:"+token, "invalid", expiration).Return(redis.NewStatusResult("", errors.New("redis error")))

		err := authService.InvalidateToken(token, expiration)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "could not invalidate token")
	})
}
