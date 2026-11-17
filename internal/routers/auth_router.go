package routers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/handlers"
)

type AuthRouter struct {
	*handlers.AuthHandler
}

func NewAuthRouter(h *handlers.AuthHandler) *AuthRouter {
	return &AuthRouter{
		AuthHandler: h,
	}
}

func (r *AuthRouter) Router(mux *http.ServeMux) {

	/*
	 * auth
	 */

	mux.HandleFunc("POST /api/auth/login", r.HandlerLogin)
	mux.HandleFunc("POST /api/auth/register", r.HandlerRegister)
	mux.HandleFunc("POST /api/auth/refresh", r.HandlerRefresh)
	mux.HandleFunc("POST /api/auth/logout", r.HandlerLogout)

}
