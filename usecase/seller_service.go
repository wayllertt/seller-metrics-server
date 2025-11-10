package usecase

import (
	"context"

	"seller-metrics-server/internal/domain"
)

type SellerService struct {
	sellers domain.SellerRepository
	metrics *MetricsService
}

func NewSellerService(sellerRepo domain.SellerRepository, metrics *MetricsService) *SellerService {
	return &SellerService{
		sellers: sellerRepo,
		metrics: metrics,
	}
}

func (s *SellerService) ListWithMetrics(ctx context.Context) ([]domain.SellerMetrics, error) {
	all, err := s.sellers.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.SellerMetrics, 0, len(all))
	for _, seller := range all {
		m, err := s.metrics.CalculateForSeller(ctx, seller.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, m)
	}
	return result, nil
}
