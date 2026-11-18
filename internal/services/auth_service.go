package services

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/benianus/gymus-go/internal/auth"
	"github.com/benianus/gymus-go/internal/config"
	"github.com/benianus/gymus-go/internal/database"
	. "github.com/benianus/gymus-go/internal/dtos"
	"github.com/benianus/gymus-go/internal/helpers"
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

func (s *AuthService) Login(req LoginRequestDto, ctx context.Context) (*AuthResponseDto, error) {
	user, err := s.GetUserByUsername(context.Background(), req.Username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if user.Username != req.Username {
		log.Println("invalid credentials!")
		return nil, errors.New("invalid credentials!")
	}

	isValid, err := auth.CheckPasswordHash(req.Password, user.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if !isValid {
		log.Println("invalid credentials!")
		return nil, errors.New("invalid credentials!")
	}

	// generate tokens
	accessToken, err := auth.MakeJwt(uuid.MustParse(user.ID.String()), s.SecretKey, time.Hour)
	refreshToken, err := auth.MakeRefreshToken()

	// save refresh token
	savedRefToken, err := s.SaveRefreshToken(ctx, database.SaveRefreshTokenParams{
		Token:     refreshToken,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		RevokedAt: sql.NullTime{},
		ExpiresAt: sql.NullTime{
			Time:  time.Now().Add(time.Hour * 24 * 7),
			Valid: true,
		},
	})

	// return response to handler
	return &AuthResponseDto{
		Username:     user.Username,
		Role:         user.Role,
		ExpiresIn:    savedRefToken.ExpiresAt.Time,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Register(registerRequestDto RegisterRequestDto, ctx context.Context) (*AuthResponseDto, error) {
	// save the password hashed
	hasehdPassword, err := auth.HashPassword(registerRequestDto.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	savedUser, err := s.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		Username:  registerRequestDto.Username,
		Password:  hasehdPassword,
		Role:      helpers.OWNER,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	accessToken, err := auth.MakeJwt(savedUser.ID, s.SecretKey, time.Hour)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	refToken, err := s.SaveRefreshToken(ctx, database.SaveRefreshTokenParams{
		Token:     refreshToken,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    savedUser.ID,
		ExpiresAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		RevokedAt: sql.NullTime{},
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &AuthResponseDto{
		Username:     savedUser.Username,
		Role:         savedUser.Role,
		ExpiresIn:    refToken.ExpiresAt.Time,
		AccessToken:  accessToken,
		RefreshToken: refToken.Token,
	}, nil
}

func (s *AuthService) Refresh(refreshRequestDto RefreshRequestDto, ctx context.Context) (*RefreshResponseDto, error) {
	user, err := s.GetUserByUsername(ctx, refreshRequestDto.Username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	refreshToken, err := s.GetRefreshTokenByUserId(ctx, user.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if refreshToken.RevokedAt.Valid || refreshToken.ExpiresAt.Time.Before(time.Now()) {
		log.Println("invalid token!")
		return nil, errors.New("ivalid token!")
	}

	accessToken, err := auth.MakeJwt(user.ID, s.SecretKey, time.Hour)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &RefreshResponseDto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken.Token,
	}, nil
}

func (s *AuthService) Logout(logoutRequestDto LogoutRequestDto, ctx context.Context) (string, error) {
	user, err := s.GetUserByUsername(ctx, logoutRequestDto.Username)
	if err != nil {
		log.Println(helpers.LOGOUT_MESSAGE)
		return helpers.LOGOUT_MESSAGE, err
	}

	refreshToken, err := s.GetRefreshTokenByUserId(ctx, user.ID)
	if err != nil {
		log.Println(helpers.LOGOUT_MESSAGE)
		return helpers.LOGOUT_MESSAGE, err
	}

	if refreshToken.RevokedAt.Valid || refreshToken.ExpiresAt.Time.Before(time.Now()) {
		log.Println(helpers.LOGOUT_MESSAGE)
		return helpers.LOGOUT_MESSAGE, errors.New(helpers.LOGOUT_MESSAGE)
	}

	if err := s.RevokeRefreshToken(ctx, refreshToken.Token); err != nil {
		log.Println(err)
		return helpers.LOGOUT_MESSAGE, err
	}

	return helpers.LOGOUT_MESSAGE, nil
}
