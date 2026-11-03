package handlers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/config"
	"github.com/benianus/gymus-go/internal/services"
)

type HandlerUser struct {
	services.AuthInterface
}

func (h *HandlerUser) HandlerLogin(w http.ResponseWriter, r *http.Request) {}

func (h *HandlerUser) HandlerRegister(w http.ResponseWriter, r *http.Request) {}

func (h *HandlerUser) HandlerRefresh(w http.ResponseWriter, r *http.Request) {}

func (h *HandlerUser) HandlerLogout(w http.ResponseWriter, r *http.Request) {}

func AuthHandlers(cfg *config.ApiConfig) HandlerUser {
	authService := services.AuthService{
		ApiConfig: cfg,
	}
	authHandler := HandlerUser{
		AuthInterface: &authService,
	}
	return authHandler
}
