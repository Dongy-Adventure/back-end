package dto

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success               bool        `json:"success"`
	Status                int         `json:"status"`
	Message               string      `json:"message"`
	Data                  interface{} `json:"data"`
	AccessToken           string      `json:"accessToken"`
	AccessTokenExpiredIn  int32       `json:"accessTokenExpiredIn"`
	RefreshToken          string      `json:"refreshToken"`
	RefreshTokenExpiredIn int32       `json:"refreshTokenExpiredIn"`
}
type LogoutRequest struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type RefreshTokenResponse struct {
	Success              bool   `json:"success"`
	Status               int    `json:"status"`
	Message              string `json:"message"`
	AccessToken          string `json:"accessToken"`
	AccessTokenExpiredIn int32  `json:"accessTokenExpiredIn"`
}
