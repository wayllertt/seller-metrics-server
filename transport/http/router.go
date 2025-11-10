package http

import "net/http"

func NewRouter(sellerHandler *SellerHandler, metricsHandler *MetricsHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/sellers/metrics", sellerHandler.ListWithMetrics)
	mux.HandleFunc("/metrics/seller", metricsHandler.GetBySellerID) // GET ?seller_id=1

	return mux
}
