package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// hash password
func HashPassword(password string) (string, error) {
	return argon2id.CreateHash(password, argon2id.DefaultParams)
}

// check hash password
func CheckPasswordHash(password, hash string) (bool, error) {
	return argon2id.ComparePasswordAndHash(password, hash)
}

// make jwt
func MakeJwt(userId uuid.UUID, tokenSecret string, expiresAt time.Duration) (string, error) {
	jwtToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    "gymus",
			Subject:   userId.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAt)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	)

	return jwtToken.SignedString([]byte(tokenSecret))
}

// validate jwt
func ValidateJwt(tokenString string, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	userId, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.UUID{}, err
	}

	return uuid.MustParse(userId), nil
}

// get Bearer Token
func GetBearerToken(header *http.Header) (string, error) {
	authHeader := header.Get("authorization")
	if len(authHeader) == 0 {
		return "", fmt.Errorf("invalid token!")
	}

	bearer := strings.Split(authHeader, " ")[0]
	token := strings.Split(authHeader, " ")[1]

	if len(bearer) == 0 && strings.ToLower(bearer) != "bearer" {
		return "", fmt.Errorf("invalid token!")
	}

	if len(token) == 0 {
		return "", fmt.Errorf("invalid token!")
	}

	return token, nil
}

// make refresh token
func MakeRefreshToken() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", nil
	}

	refreshToken := hex.EncodeToString(key)
	return refreshToken, nil
}
