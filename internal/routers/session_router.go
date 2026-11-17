package routers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/handlers"
)

type SessionRouter struct {
	*handlers.SessionHandler
}

func NewSessionRouter(h *handlers.SessionHandler) *SessionRouter {
	return &SessionRouter{
		SessionHandler: h,
	}
}

func (r *SessionRouter) Router(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/sessions", r.GetAllSessions)
	mux.HandleFunc("POST /api/sessions", r.CreateSession)
}
