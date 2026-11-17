package handlers

import "github.com/benianus/gymus-go/internal/services"

type StoreHandler struct {
	*services.StoreService
}

func NewStoreHandler(s *services.StoreService) *StoreHandler {
	return &StoreHandler{
		StoreService: s,
	}
}
