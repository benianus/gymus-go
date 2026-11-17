package routers

import (
	"net/http"

	"github.com/benianus/gymus-go/internal/handlers"
)

type StoreRouter struct {
	*handlers.StoreHandler
}

func NewStoreRouter(h *handlers.StoreHandler) *StoreRouter {
	return &StoreRouter{
		StoreHandler: h,
	}
}

func (r *StoreRouter) Route(mux *http.ServeMux) {

}
