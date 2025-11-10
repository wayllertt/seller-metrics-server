package usecase

import (
	"context"
	"time"

	"seller-metrics-server/internal/domain"
)

type OrderService struct {
	orders  domain.OrderRepository
	sellers domain.SellerRepository
}

func NewOrderService(orderRepo domain.OrderRepository, sellerRepo domain.SellerRepository) *OrderService {
	return &OrderService{
		orders:  orderRepo,
		sellers: sellerRepo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, sellerID int, promisedAfterHours int) error {
	_, err := s.sellers.GetByID(ctx, sellerID)
	if err != nil {
		return err
	}

	now := time.Now()

	order := domain.Order{
		SellerID:   sellerID,
		CreatedAt:  now,
		PromisedAt: now.Add(time.Duration(promisedAfterHours) * time.Hour),
	}

	return s.orders.Add(ctx, order)
}

func (s *OrderService) MarkDelivered(ctx context.Context, orderID int) error {
	return nil
}

func (s *OrderService) CreateDemoOrders(ctx context.Context) error {
	now := time.Now()

	o1 := domain.Order{
		SellerID:   1,
		CreatedAt:  now.Add(-30 * time.Hour),
		PromisedAt: now.Add(-6 * time.Hour),
	}
	d1 := now.Add(-8 * time.Hour)
	o1.DeliveredAt = &d1

	o2 := domain.Order{
		SellerID:   1,
		CreatedAt:  now.Add(-40 * time.Hour),
		PromisedAt: now.Add(-10 * time.Hour),
	}
	d2 := now.Add(-9 * time.Hour)
	o2.DeliveredAt = &d2

	o3 := domain.Order{
		SellerID:   2,
		CreatedAt:  now.Add(-20 * time.Hour),
		PromisedAt: now.Add(-5 * time.Hour),
	}
	d3 := now.Add(-6 * time.Hour)
	o3.DeliveredAt = &d3

	if err := s.orders.Add(ctx, o1); err != nil {
		return err
	}
	if err := s.orders.Add(ctx, o2); err != nil {
		return err
	}
	if err := s.orders.Add(ctx, o3); err != nil {
		return err
	}

	return nil
}
