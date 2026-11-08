package services

import (
	"github.com/benianus/gymus-go/internal/config"
)

type MembershipService struct {
	*config.ApiConfig
}

func NewMembershipService(cfg *config.ApiConfig) *MembershipService {
	return &MembershipService{
		ApiConfig: cfg,
	}
}
