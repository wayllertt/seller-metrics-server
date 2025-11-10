package app

import (
	"net/http"

	mem "seller-metrics/internal/storage/memory"
	thttp "seller-metrics/internal/transport/http"
	"seller-metrics/internal/usecase"
)

func NewHTTPServer() http.Handler {
	sellerRepo := mem.NewSellerRepoMemory()
	orderRepo := mem.NewOrderRepoMemory()

	metricsService := usecase.NewMetricsService(sellerRepo, orderRepo)
	sellerService := usecase.NewSellerService(sellerRepo, metricsService)

	sellerHandler := thttp.NewSellerHandler(sellerService)
	metricsHandler := thttp.NewMetricsHandler(metricsService)

	return thttp.NewRouter(sellerHandler, metricsHandler)
}
