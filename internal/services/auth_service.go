package services

import (
	"context"
	"time"
	"uuid"

	. "github.com/benianus/gymus-go/internal/config"
)

type AuthService struct {
	*ApiConfig
}

func NewAuthService(cfg *ApiConfig) *AuthService {
	return &AuthService{
		ApiConfig: cfg,
	}
}

type LoginRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Username     string    `json:"username"`
	Role         string    `json:"role"`
	ExpiresIn    time.Time `json:"expires_in"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type RegisterRequestDto struct {
	Username string `json:"username"`
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

func (auth *AuthService) Login(loginRequestDto LoginRequestDto) (*AuthResponseDto, error) {
	user, err := auth.GetUserByUsername(context.Background(), loginRequestDto.Username)
	if err != nil {
		return nil, err
	}

	return &AuthResponseDto{
		Username:     user.Username,
		Role:         user.Role,
		ExpiresIn:    time.Now().Add(time.Hour),
		AccessToken:  uuid.NewV7().String(),
		RefreshToken: uuid.NewV7().String(),
	}, nil
}

func (auth *AuthService) Register(registerRequestDto RegisterRequestDto) (AuthResponseDto, error) {
	return AuthResponseDto{}, nil
}

func (auth *AuthService) Refresh(refreshRequestDto RefreshRequestDto) (RefreshResponseDto, error) {
	return RefreshResponseDto{}, nil
}

func (auth *AuthService) Logout(logoutRequestDto LogoutRequestDto) error {
	return nil
}
