package dtos

import "time"

type LoginRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Username     string    `json:"username" validate:"required"`
	Role         string    `json:"role"`
	ExpiresIn    time.Time `json:"expires_in"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type RegisterRequestDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
}
type RefreshRequestDto struct {
	Username     string `json:"username"`
	RefreshToken string `json:"refresh_token"`
}
type LogoutRequestDto struct {
	Username     string `json:"username"`
	RefreshToken string `json:"refresh_token"`
}
type RefreshResponseDto struct {
	Username     string `json:"username"`
	RefreshToken string `json:"refresh_token"`
}
