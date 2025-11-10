package transporthttp

import stdhttp "net/http"

func NewRouter(sellerHandler *SellerHandler, metricsHandler *MetricsHandler) stdhttp.Handler {
	mux := stdhttp.NewServeMux()

	mux.HandleFunc("/sellers/metrics", sellerHandler.ListWithMetrics)
	mux.HandleFunc("/metrics/seller", metricsHandler.GetBySellerID) // GET ?seller_id=1

	return mux
}
