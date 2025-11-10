package memory

import (
	"context"
	"errors"
	"sync"
	"time"

	"seller-metrics/internal/domain"
)

type SellerRepoMemory struct {
	mu      sync.RWMutex
	storage map[int]domain.Seller
}

func NewSellerRepoMemory() *SellerRepoMemory {
	now := time.Now()
	return &SellerRepoMemory{
		storage: map[int]domain.Seller{
			1: {ID: 1, Name: "Seller One", CreatedAt: now},
			2: {ID: 2, Name: "Seller Two", CreatedAt: now},
		},
	}
}

func (r *SellerRepoMemory) GetAll(ctx context.Context) ([]domain.Seller, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]domain.Seller, 0, len(r.storage))
	for _, s := range r.storage {
		result = append(result, s)
	}
	return result, nil
}

func (r *SellerRepoMemory) GetByID(ctx context.Context, id int) (*domain.Seller, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	s, ok := r.storage[id]
	if !ok {
		return nil, errors.New("seller not found")
	}
	return &s, nil
}

func (r *SellerRepoMemory) Update(ctx context.Context, s domain.Seller) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.storage[s.ID] = s
	return nil
}
