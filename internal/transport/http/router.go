package transporthttp

import "net/http"

func NewRouter(
	sellerHandler *SellerHandler,
	metricsHandler *MetricsHandler,
	orderHandler *OrderHandler,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/sellers/metrics", sellerHandler.ListWithMetrics)
	mux.HandleFunc("/metrics/seller", metricsHandler.GetBySellerID)

	mux.HandleFunc("/orders/demo", orderHandler.CreateDemoOrders)
	mux.HandleFunc("/orders", orderHandler.CreateOrder)

	return mux
}
