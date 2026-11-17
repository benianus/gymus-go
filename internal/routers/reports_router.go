package routers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/handlers"
)

type ReportsRouter struct {
	*handlers.ReportsHandler
}

func NewReportsRouter(h *handlers.ReportsHandler) *ReportsRouter {
	return &ReportsRouter{
		ReportsHandler: h,
	}
}

func (r *ReportsRouter) Router(mux *http.ServeMux) {
    
}
