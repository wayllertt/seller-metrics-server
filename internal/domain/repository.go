package domain

import "context"

type SellerRepository interface {
	GetAll(ctx context.Context) ([]Seller, error)
	GetByID(ctx context.Context, id int) (*Seller, error)
	Update(ctx context.Context, s Seller) error
}

type OrderRepository interface {
	GetBySellerID(ctx context.Context, sellerID int) ([]Order, error)
	Add(ctx context.Context, o Order) error
}
