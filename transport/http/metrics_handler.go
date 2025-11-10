package http

import (
	"net/http"
	"strconv"

	"seller-metrics/internal/usecase"
)

type MetricsHandler struct {
	metrics *usecase.MetricsService
}

func NewMetricsHandler(m *usecase.MetricsService) *MetricsHandler {
	return &MetricsHandler{metrics: m}
}

func (h *MetricsHandler) GetBySellerID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.URL.Query().Get("seller_id")
	if idStr == "" {
		http.Error(w, "seller_id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid seller_id", http.StatusBadRequest)
		return
	}

	metrics, err := h.metrics.CalculateForSeller(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, metrics)
}
