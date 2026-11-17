package routers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/config"
	"github.com/benianus/gymus-go/internal/handlers"
	"github.com/benianus/gymus-go/internal/services"
)

func Router(cfg *config.ApiConfig, mux *http.ServeMux) {
	authService := services.NewAuthService(cfg)
	authHandler := handlers.NewAuthHandler(authService)
	authRouter := NewAuthRouter(authHandler)

	authRouter.Router(mux)

	// Memberships
	mebershipService := services.NewMembershipService(cfg)
	membershipHandler := handlers.NewMembershipHandler(mebershipService)
	membershipRouter := NewMembershipRouter(membershipHandler)

	membershipRouter.Router(mux)

	// sessions
	sessionService := services.NewSessionService(cfg)
	sessionHandler := handlers.NewSessionHandler(sessionService)
	sessionRouter := NewSessionRouter(sessionHandler)

	sessionRouter.Router(mux)

	// store
	storeService := services.NewStoreService(cfg)
	storeHandler := handlers.NewStoreHandler(storeService)
	storeRouter := NewStoreRouter(storeHandler)

	storeRouter.Route(mux)

	// reports
	reportService := services.NewReportService(cfg)
	reportHandler := handlers.NewReportHandler(reportService)
	reportRouter := NewReportsRouter(reportHandler)

	reportRouter.Router(mux)
}
