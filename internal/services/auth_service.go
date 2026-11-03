package services

import (
	"time"

	"github.com/benianus/gymus-go/internal/config"
)

type AuthService struct {
	*config.ApiConfig
}

type LoginRequestDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Username     string
	Role         string
	ExpiresIn    time.Time
	AccessToken  string
	RefreshToken string
}

type RegisterRequestDto struct{}
type RefreshRequestDto struct{}
type LogoutRequestDto struct{}
type RefreshResponseDto struct{}

type AuthInterface interface {
	Login(LoginRequestDto) (AuthResponseDto, error)
	Register(RegisterRequestDto) (AuthResponseDto, error)
	Refresh(RefreshRequestDto) (RefreshResponseDto, error)
	Logout(LogoutRequestDto) error
}

func (auth *AuthService) Login(loginRequestDto LoginRequestDto) (AuthResponseDto, error) {
	return AuthResponseDto{}, nil
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
