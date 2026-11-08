package routers

import "github.com/benianus/gymus-go/internal/handlers"

type ReportsRouter struct {
	*handlers.ReportsHandler
}

func NewReportsRouter(handler *handlers.ReportsHandler) *ReportsRouter {
	return &ReportsRouter{
		ReportsHandler: handler,
	}
}
