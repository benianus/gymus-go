package services

import "github.com/benianus/gymus-go/internal/config"

type ReportsService struct {
	*config.ApiConfig
}

func NewReportService(cfg *config.ApiConfig) *ReportsService {
	return &ReportsService{
		ApiConfig: cfg,
	}
}
