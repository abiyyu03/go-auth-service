package dto

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	UserData  UserResponse  `json:"user_data"`
	TokenData TokenResponse `json:"token_data"`
}

type TokenResponse struct {
	AccessToken    string `json:"access_token"`
	AccessExpires  int64  `json:"access_expired_at"`
	RefreshToken   string `json:"refresh_token"`
	RefreshExpires int64  `json:"refresh_expired_at"`
}
