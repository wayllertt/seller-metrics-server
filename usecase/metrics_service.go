package usecase

import (
	"context"
	"math"

	"seller-metrics-server/internal/domain"
)

type MetricsService struct {
	sellers domain.SellerRepository
	orders  domain.OrderRepository
}

func NewMetricsService(sellerRepo domain.SellerRepository, orderRepo domain.OrderRepository) *MetricsService {
	return &MetricsService{
		sellers: sellerRepo,
		orders:  orderRepo,
	}
}

func (m *MetricsService) CalculateForSeller(ctx context.Context, sellerID int) (domain.SellerMetrics, error) {
	orders, err := m.orders.GetBySellerID(ctx, sellerID)
	if err != nil {
		return domain.SellerMetrics{}, err
	}

	var (
		totalHours float64
		count      int
		onTime     int
	)

	for _, o := range orders {
		if o.DeliveredAt == nil {
			continue
		}

		hours := o.DeliveredAt.Sub(o.CreatedAt).Hours()
		totalHours += hours
		count++

		if !o.DeliveredAt.After(o.PromisedAt) {
			onTime++
		}
	}

	var avgHours, onTimeRate float64
	if count > 0 {
		avgHours = totalHours / float64(count)
		onTimeRate = float64(onTime) / float64(count)
	}

	isBlocked := onTimeRate < 0.8 && count >= 3

	return domain.SellerMetrics{
		SellerID:         sellerID,
		AvgDeliveryHours: round2(avgHours),
		OnTimeRate:       round2(onTimeRate),
		IsBlocked:        isBlocked,
	}, nil
}

func round2(v float64) float64 {
	return math.Round(v*100) / 100
}
