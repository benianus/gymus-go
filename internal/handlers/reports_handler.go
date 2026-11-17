package handlers

import "github.com/benianus/gymus-go/internal/services"

type ReportsHandler struct {
	*services.ReportsService
}

func NewReportHandler(s *services.ReportsService) *ReportsHandler {
	return &ReportsHandler{
		ReportsService: s,
	}
}
