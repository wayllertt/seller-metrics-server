package app

import (
	"net/http"

	thttp "seller-metrics-server/internal/transport/http"
	"seller-metrics-server/internal/usecase"
	mem "seller-metrics-server/storage/memory"
)

func NewHTTPServer() http.Handler {
	sellerRepo := mem.NewSellerRepoMemory()
	orderRepo := mem.NewOrderRepoMemory()

	metricsService := usecase.NewMetricsService(sellerRepo, orderRepo)
	sellerService := usecase.NewSellerService(sellerRepo, metricsService)
	orderService := usecase.NewOrderService(orderRepo, sellerRepo)

	sellerHandler := thttp.NewSellerHandler(sellerService)
	metricsHandler := thttp.NewMetricsHandler(metricsService)
	orderHandler := thttp.NewOrderHandler(orderService)

	return thttp.NewRouter(sellerHandler, metricsHandler, orderHandler)
}
