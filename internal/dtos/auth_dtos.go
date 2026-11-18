package dtos

import "time"

type LoginRequestDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequestDto struct {
	Username string `json:"username" validate:"required,min=3,max=255"`
	Password string `json:"password" validate:"required,min=8"`
}
type RefreshRequestDto struct {
	Username     string `json:"username" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}
type LogoutRequestDto struct {
	Username     string `json:"username" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}
type RefreshResponseDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthResponseDto struct {
	Username     string    `json:"username"`
	Role         string    `json:"role"`
	ExpiresIn    time.Time `json:"expires_in"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}
