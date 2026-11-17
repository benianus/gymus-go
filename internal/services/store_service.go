package services

import "github.com/benianus/gymus-go/internal/config"

type StoreService struct {
	*config.ApiConfig
}

func NewStoreService(cfg *config.ApiConfig) *StoreService {
	return &StoreService{
		ApiConfig: cfg,
	}
}

func (s *StoreService) GetAllProducts() {}

func (s *StoreService) GetProductById() {}

func (s *StoreService) CreateProduct() {}

func (s *StoreService) CreateSale() {}
