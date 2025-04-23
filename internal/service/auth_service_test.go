package service_test

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Dongy-s-Advanture/back-end/internal/config"
	"github.com/Dongy-s-Advanture/back-end/internal/dto"
	"github.com/Dongy-s-Advanture/back-end/internal/enum/tokenmode"
	"github.com/Dongy-s-Advanture/back-end/internal/model"
	"github.com/Dongy-s-Advanture/back-end/internal/service"
	mock "github.com/Dongy-s-Advanture/back-end/pkg/mock/repository"
	"github.com/Dongy-s-Advanture/back-end/pkg/utils/token"
	"github.com/gin-gonic/gin"
	redismock "github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestSellerLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSellerRepo := mock.NewMockISellerRepository(ctrl)

	// mock Redis (not used in login, but needed for constructor)
	db, _ := redismock.NewClientMock()

	// config
	cfg := &config.Config{
		Auth: config.AuthConfig{
			AccessTokenLifespanMinutes:  15,
			RefreshTokenLifespanMinutes: 60,
			AccessTokenSecret:           "testaccesssecret",
			RefreshTokenSecret:          "testrefreshsecret",
		},
	}

	authService := service.NewAuthService(cfg, db, mockSellerRepo, nil)

	// Prepare test data
	sellerID := primitive.NewObjectID()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	mockModel := &model.Seller{
		SellerID: sellerID,
		Username: "seller",
		Password: string(hashedPassword),
	}

	req := &dto.LoginRequest{
		Username: "seller",
		Password: "password",
	}

	mockSellerRepo.EXPECT().
		GetSellerByUsername(req).
		Return(mockModel, nil).
		Times(1)

	sellerDTO, accessToken, refreshToken, err := authService.SellerLogin(req)

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	assert.NotEmpty(t, refreshToken)
	assert.Equal(t, mockModel.Username, sellerDTO.Username)
	assert.Equal(t, mockModel.SellerID, sellerDTO.SellerID)
}
func TestBuyerLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBuyerRepo := mock.NewMockIBuyerRepository(ctrl)
	db, _ := redismock.NewClientMock()

	cfg := &config.Config{
		Auth: config.AuthConfig{
			AccessTokenLifespanMinutes:  15,
			RefreshTokenLifespanMinutes: 60,
			AccessTokenSecret:           "testaccesssecret",
			RefreshTokenSecret:          "testrefreshsecret",
		},
	}

	authService := service.NewAuthService(cfg, db, nil, mockBuyerRepo)

	buyerID := primitive.NewObjectID()

	buyerModel := &model.Buyer{
		BuyerID:  buyerID,
		Username: "buyer",
		Password: "$2a$10$saltrandomizedhashpassword",
	}

	req := &dto.LoginRequest{
		Username: "buyer",
		Password: "any-password",
	}

	mockBuyerRepo.EXPECT().GetBuyerByUsername(req).Return(buyerModel, nil).Times(1)

	// override bcrypt to avoid error (or use hash like above)
	authService = service.NewAuthService(cfg, db, nil, mockBuyerRepo)

	buyerDTO, accessToken, refreshToken, err := authService.BuyerLogin(req)

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	assert.NotEmpty(t, refreshToken)
	assert.Equal(t, buyerModel.BuyerID, buyerDTO.BuyerID)
}

func TestLogout_Success(t *testing.T) {
	db, mock := redismock.NewClientMock()
	cfg := &config.Config{
		Auth: config.AuthConfig{
			AccessTokenLifespanMinutes:  10,
			RefreshTokenLifespanMinutes: 20,
		},
	}

	authService := service.NewAuthService(cfg, db, nil, nil)

	mock.ExpectSetEx("blacklist:access-token", "invalid", 10*time.Minute).SetVal("OK")
	mock.ExpectSetEx("blacklist:refresh-token", "invalid", 20*time.Minute).SetVal("OK")

	err := authService.Logout("access-token", "refresh-token")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRefreshToken_Success(t *testing.T) {
	db, rdmock := redismock.NewClientMock()

	cfg := &config.Config{
		Auth: config.AuthConfig{
			AccessTokenLifespanMinutes:  10,
			RefreshTokenLifespanMinutes: 20,
			AccessTokenSecret:           "testaccess",
			RefreshTokenSecret:          "testrefresh",
		},
	}

	authService := service.NewAuthService(cfg, db, nil, nil)

	// simulate a user ID
	userID := primitive.NewObjectID().Hex()

	// manually generate a valid refresh token
	refreshToken, _ := token.GenerateToken(cfg, userID, tokenmode.REFRESH_TOKEN)
	rdmock.ExpectExists("blacklist:" + refreshToken).SetVal(0)

	// use gin context with header set
	c, _ := gin.CreateTestContext(nil)
	c.Request = httptest.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Authorization", "Bearer "+refreshToken)

	accessToken, err := authService.RefreshToken(c)

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)

	assert.NoError(t, rdmock.ExpectationsWereMet())

}
