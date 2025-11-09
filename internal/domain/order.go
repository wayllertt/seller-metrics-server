package domain

import "time"

type Order struct {
	ID          int
	SellerID    int
	CreatedAt   time.Time
	DeliveredAt *time.Time
	PromisedAt  time.Time
}
