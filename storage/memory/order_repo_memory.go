package memory

import (
	"context"
	"sync"

	"seller-metrics-server/internal/domain"
)

type OrderRepoMemory struct {
	mu     sync.RWMutex
	orders []domain.Order
	nextID int
}

func NewOrderRepoMemory() *OrderRepoMemory {
	return &OrderRepoMemory{
		orders: make([]domain.Order, 0),
		nextID: 1,
	}
}

func (r *OrderRepoMemory) GetBySellerID(ctx context.Context, sellerID int) ([]domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var res []domain.Order
	for _, o := range r.orders {
		if o.SellerID == sellerID {
			res = append(res, o)
		}
	}
	return res, nil
}

func (r *OrderRepoMemory) Add(ctx context.Context, o domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	o.ID = r.nextID
	r.nextID++
	r.orders = append(r.orders, o)
	return nil
}
