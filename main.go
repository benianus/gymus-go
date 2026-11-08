package main

import (
	"log"
	"net/http"

	_ "github.com/jackc/pgx"

	"github.com/benianus/gymus-go/internal/config"
	"github.com/benianus/gymus-go/internal/handlers"
	"github.com/benianus/gymus-go/internal/routers"
	"github.com/benianus/gymus-go/internal/services"
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

	/**
	* initilize server
	* routing & handlers
	 */
	mux := http.NewServeMux()

	// Auth
	authService := services.NewAuthService(&cfg)
	authHandler := handlers.NewAuthHandler(authService)
	authRouter := routers.NewAuthRouter(authHandler)

	authRouter.Router(mux)

	// Memberships
	mebershipService := services.NewMembershipService(&cfg)
	membershipHandler := handlers.NewMembershipHandler(mebershipService)
	membershipRouter := routers.NewMembershipRouter(membershipHandler)

	membershipRouter.Router(mux)
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
