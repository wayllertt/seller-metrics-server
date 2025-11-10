package domain

import "context"

type SellerRepository struct {
	GetAll (ctx context.Context) ([]Seller, error)
	GetById (ctx context.Context, id int) (*Seller, error)
	Update (tx context.Context, s Seller) (error) 
}

type OrderRepository struct {
	GetBySellerID (ctx context.Context, sellerID int) ([]Order, error)
	Add (ctx context.Context, o Order) error
}