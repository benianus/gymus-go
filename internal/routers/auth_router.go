package routers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/handlers"
)

type AuthRouter struct {
	*handlers.AuthHandler
}

func NewAuthRouter(authHandler *handlers.AuthHandler) *AuthRouter {
	return &AuthRouter{
		AuthHandler: authHandler,
	}
}

func (route *AuthRouter) Router(mux *http.ServeMux) {
	
	// membershipHandlers := MembershipHandlers(cfg)
	/*
	 * auth
	 */
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	mux.HandleFunc("POST /api/auth/login", route.HandlerLogin)
	mux.HandleFunc("POST /api/auth/register", route.HandlerRegister)
	mux.HandleFunc("POST /api/auth/refresh", route.HandlerRefresh)
	mux.HandleFunc("POST /api/auth/logout", route.HandlerLogout)
	/**
	 * store
	 */

	/**
	 * reports
	 */

}
