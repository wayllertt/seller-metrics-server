package usecase

import (
	"context"
	"testing"
	"time"

	"seller-metrics-server/internal/domain"
)

type fakeOrderRepo struct {
	orders []domain.Order
}

func (f *fakeOrderRepo) GetBySellerID(ctx context.Context, sellerID int) ([]domain.Order, error) {
	var res []domain.Order
	for _, o := range f.orders {
		if o.SellerID == sellerID {
			res = append(res, o)
		}
	}
	return res, nil
}
func (f *fakeOrderRepo) Add(ctx context.Context, o domain.Order) error {
	f.orders = append(f.orders, o)
	return nil
}

type fakeSellerRepo struct{}

func (f *fakeSellerRepo) GetAll(ctx context.Context) ([]domain.Seller, error) { return nil, nil }
func (f *fakeSellerRepo) GetByID(ctx context.Context, id int) (*domain.Seller, error) {
	return &domain.Seller{ID: id, Name: "Test"}, nil
}
func (f *fakeSellerRepo) Update(ctx context.Context, s domain.Seller) error { return nil }

func TestMetricsService_CalculateForSeller(t *testing.T) {
	now := time.Now()

	d1 := now.Add(10 * time.Hour)
	d2 := now.Add(20 * time.Hour)
	d3 := now.Add(40 * time.Hour)

	fakeOrders := &fakeOrderRepo{
		orders: []domain.Order{
			{
				SellerID:    1,
				CreatedAt:   now,
				PromisedAt:  now.Add(24 * time.Hour),
				DeliveredAt: &d1, // вовремя
			},
			{
				SellerID:    1,
				CreatedAt:   now,
				PromisedAt:  now.Add(24 * time.Hour),
				DeliveredAt: &d2, // вовремя
			},
			{
				SellerID:    1,
				CreatedAt:   now,
				PromisedAt:  now.Add(24 * time.Hour),
				DeliveredAt: &d3, // опоздал
			},
		},
	}

	ms := NewMetricsService(&fakeSellerRepo{}, fakeOrders)

	metrics, err := ms.CalculateForSeller(context.Background(), 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if metrics.OnTimeRate < 0.66 || metrics.OnTimeRate > 0.68 {
		t.Errorf("expected onTimeRate around 0.67, got %v", metrics.OnTimeRate)
	}

	if !metrics.IsBlocked {
		t.Errorf("expected seller to be blocked, got not blocked")
	}
}
