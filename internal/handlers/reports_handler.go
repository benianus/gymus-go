package handlers

import "github.com/benianus/gymus-go/internal/services"

type ReportsHandler struct {
	*services.ReportsService
}

func NewReportsService(service *services.ReportsService) *ReportsHandler {
	return &ReportsHandler{
		ReportsService: service,
	}
}
