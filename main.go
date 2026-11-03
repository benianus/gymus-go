package main

import (
	"log"
	"net/http"

	_ "github.com/jackc/pgx"

	"github.com/benianus/gymus-go/internal/config"
	"github.com/benianus/gymus-go/internal/handlers"
)

func main() {
	cfg := config.ApiConfig{}

	/**
	 * load environement variables
	 */
	if err := cfg.LoadEnvVariables(); err != nil {
		log.Fatal(err)
		return
	}

	/**
	 * initilize database
	 */
	if err := cfg.LoadDatabase(); err != nil {
		log.Fatal(err)
		return
	}

	/*
	 * handlers
	 */
	authHandlers := handlers.AuthHandlers(&cfg)
	membershipHandlers := handlers.MembershipHandlers(&cfg)
	/**
	* initilize server
	* routing & handlers
	 */
	mux := http.NewServeMux()

	/*
	 * auth
	 */
	mux.HandleFunc("POST /api/auth/login", authHandlers.HandlerLogin)
	mux.HandleFunc("POST /api/auth/register", authHandlers.HandlerRegister)
	mux.HandleFunc("POST /api/auth/refresh", authHandlers.HandlerRefresh)
	mux.HandleFunc("POST /api/auth/logout", authHandlers.HandlerLogout)

	/**
	 * membership
	 */
	mux.HandleFunc("GET /api/members", membershipHandlers.GetAllMembers)
	mux.HandleFunc("GET /api/members/{memberId}", membershipHandlers.GetMemberCard)
	mux.HandleFunc("POST /api/members", membershipHandlers.RegisterMemberhsip)
	mux.HandleFunc("POST /api/members/{memberId}/renew", membershipHandlers.RenewMembership)
	mux.HandleFunc("POST /api/members/record/{memberId}/attendance", membershipHandlers.RecordAttandance)
	mux.HandleFunc("PUT /api/members/{memberId}", membershipHandlers.UpdateMember)
	mux.HandleFunc("DELETE /api/members/{memberId}", membershipHandlers.DeleteMember)
	mux.HandleFunc("GET /api/members/recent-checks", membershipHandlers.GetRecentChecksInMembers)

	/**
	 * session
	 */

	/**
	 * store
	 */

	/**
	 * reports
	 */
	
	/**
	 * initialize server
	 */
	server := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	/*
	 * log server informations
	 */
	log.Printf("Serving files from  on port: %s\n", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
