package domain

type SellerMetrics struct {
	SellerID         int
	AvgDeliveryHours float64
	OnTimeRate       float64
	IsBlocked        bool
}
