package domain

type SellerMetrics struct {
	SellerID      int
	DeliveryHours float64
	onTimeRate    float64
	isBlocked     bool
}
