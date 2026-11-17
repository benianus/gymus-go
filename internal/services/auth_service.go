package services

import (
	"context"
	"log"
	"time"

	"github.com/benianus/gymus-go/internal/auth"
	"github.com/benianus/gymus-go/internal/config"
	"github.com/benianus/gymus-go/internal/database"
	. "github.com/benianus/gymus-go/internal/dtos"
	"github.com/google/uuid"
)

type AuthService struct {
	*config.ApiConfig
}

func NewAuthService(cfg *config.ApiConfig) *AuthService {
	return &AuthService{
		ApiConfig: cfg,
	}
}

func (s *AuthService) Login(req LoginRequestDto) (*AuthResponseDto, error) {
	user, err := s.GetUserByUsername(context.Background(), req.Username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	isValid, err := auth.CheckPasswordHash(req.Password, user.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if !isValid {
		log.Println(err)
		return nil, err
	}

	accessToken, err := auth.MakeJwt(uuid.MustParse(user.ID.String()), s.SecretKey, time.Hour)
	refreshToken, err := auth.MakeRefreshToken()

	return &AuthResponseDto{
		Username:     user.Username,
		Role:         user.Role,
		ExpiresIn:    time.Now().Add(time.Hour),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Register(registerRequestDto RegisterRequestDto, ctx context.Context) (*AuthResponseDto, error) {
	// save the password hashed
	hasehdPassword, err := auth.HashPassword(registerRequestDto.Password)
	if err != nil {
		return nil, err
	}
	newUser, err := s.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		Username:  registerRequestDto.Username,
		Password:  hasehdPassword,
		Role:      "OWNER",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	accessToken, err := auth.MakeJwt(newUser.ID, s.SecretKey, time.Hour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		return nil, err
	}

	return &AuthResponseDto{
		Username:     newUser.Username,
		Role:         newUser.Password,
		ExpiresIn:    time.Now().Add(7),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Refresh(refreshRequestDto RefreshRequestDto) (RefreshResponseDto, error) {
	return RefreshResponseDto{}, nil
}

func (s *AuthService) Logout(logoutRequestDto LogoutRequestDto) error {
	return nil
}
