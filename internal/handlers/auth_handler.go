package handlers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/services"
)

type AuthHandler struct {
	*services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: service,
	}
}

func (handler *AuthHandler) HandlerLogin(w http.ResponseWriter, r *http.Request) {
    
}

func (handler *AuthHandler) HandlerRegister(w http.ResponseWriter, r *http.Request) {}

func (handler *AuthHandler) HandlerRefresh(w http.ResponseWriter, r *http.Request) {}

func (handler *AuthHandler) HandlerLogout(w http.ResponseWriter, r *http.Request) {}


