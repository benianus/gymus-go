package handlers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/services"
)

type MembershipHandler struct {
	*services.MembershipService
}

func NewMembershipHandler(s *services.MembershipService) *MembershipHandler {
	return &MembershipHandler{
		MembershipService: s,
	}
}

func (h *MembershipHandler) GetAllMembers(w http.ResponseWriter, r *http.Request) {}

func (h *MembershipHandler) GetMemberCard(w http.ResponseWriter, r *http.Request) {}

func (h *MembershipHandler) RegisterMemberhsip(w http.ResponseWriter, r *http.Request) {}

func (h *MembershipHandler) RenewMembership(w http.ResponseWriter, r *http.Request) {}

func (h *MembershipHandler) RecordAttandance(w http.ResponseWriter, r *http.Request) {}

func (h *MembershipHandler) UpdateMember(w http.ResponseWriter, r *http.Request) {}

func (h *MembershipHandler) DeleteMember(w http.ResponseWriter, r *http.Request) {}

func (h *MembershipHandler) GetRecentChecksInMembers(w http.ResponseWriter, r *http.Request) {}
