package routers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/handlers"
)

type MembershipRouter struct {
	*handlers.MembershipHandler
}

func NewMembershipRouter(handler *handlers.MembershipHandler) *MembershipRouter {
	return &MembershipRouter{
		MembershipHandler: handler,
	}
}

func (r *MembershipRouter) Router(mux *http.ServeMux) {
	/**
	 * membership
	 */
	mux.HandleFunc("GET /api/members", r.GetAllMembers)
	mux.HandleFunc("GET /api/members/{memberId}", r.GetMemberCard)
	mux.HandleFunc("POST /api/members", r.RegisterMemberhsip)
	mux.HandleFunc("POST /api/members/{memberId}/renew", r.RenewMembership)
	mux.HandleFunc("POST /api/members/record/{memberId}/attendance", r.RecordAttandance)
	mux.HandleFunc("PUT /api/members/{memberId}", r.UpdateMember)
	mux.HandleFunc("DELETE /api/members/{memberId}", r.DeleteMember)
	mux.HandleFunc("GET /api/members/recent-checks", r.GetRecentChecksInMembers)
}
