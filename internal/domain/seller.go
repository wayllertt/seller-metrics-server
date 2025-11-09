package domain

import "time"

type Seller struct {
	ID        int
	Name      string
	IsBlocked bool
	CreatedAt time.Time
}
