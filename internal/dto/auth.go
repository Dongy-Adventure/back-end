package dto

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success      bool        `json:"success"`
	Status       int         `json:"status"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
}
