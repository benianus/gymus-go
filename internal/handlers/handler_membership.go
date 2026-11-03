package handlers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/config"
	"github.com/benianus/gymus-go/internal/services"
)

type HandlerMembership struct {
	services.MembershipService
}

func (h *HandlerMembership) GetAllMembers(w http.ResponseWriter, r *http.Request) {}

func (h *HandlerMembership) GetMemberCard(w http.ResponseWriter, r *http.Request) {}

func (h *HandlerMembership) RegisterMemberhsip(w http.ResponseWriter, r *http.Request) {}

func (h *HandlerMembership) RenewMembership(w http.ResponseWriter, r *http.Request) {}

func (h *HandlerMembership) RecordAttandance(w http.ResponseWriter, r *http.Request) {}

func (h *HandlerMembership) UpdateMember(w http.ResponseWriter, r *http.Request) {}

func (h *HandlerMembership) DeleteMember(w http.ResponseWriter, r *http.Request) {}

func (h *HandlerMembership) GetRecentChecksInMembers(w http.ResponseWriter, r *http.Request) {}

func MembershipHandlers(cfg *config.ApiConfig) HandlerMembership {
	membershipService := services.MembershipService{
		ApiConfig: cfg,
	}

	membershipHandlers := HandlerMembership{
		MembershipService: membershipService,
	}

	return membershipHandlers
}
