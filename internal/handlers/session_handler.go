package handlers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/services"
)

type SessionHandler struct {
	*services.SessionService
}

func NewSessionHandler(s *services.SessionService) *SessionHandler {
	return &SessionHandler{
		SessionService: s,
	}
}

func (h *SessionHandler) GetAllSessions(w http.ResponseWriter, r *http.Request) {}

func (h *SessionHandler) CreateSession(w http.ResponseWriter, r *http.Request) {}
