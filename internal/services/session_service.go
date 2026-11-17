package services

import "github.com/benianus/gymus-go/internal/config"

type SessionService struct {
	*config.ApiConfig
}

func NewSessionService(cfg *config.ApiConfig) *SessionService {
	return &SessionService{
		ApiConfig: cfg,
	}
}

func (s *SessionService) GetAllSessions() {}

func (s *SessionService) CreateSession() {}
