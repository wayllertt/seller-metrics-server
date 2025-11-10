package transporthttp

import (
	stdhttp "net/http"
	"strconv"

	"seller-metrics-server/usecase"
)

type MetricsHandler struct {
	metrics *usecase.MetricsService
}

func NewMetricsHandler(m *usecase.MetricsService) *MetricsHandler {
	return &MetricsHandler{metrics: m}
}

func (h *MetricsHandler) GetBySellerID(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()
	idStr := r.URL.Query().Get("seller_id")
	if idStr == "" {
		stdhttp.Error(w, "seller_id is required", stdhttp.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		stdhttp.Error(w, "invalid seller_id", stdhttp.StatusBadRequest)
		return
	}

	metrics, err := h.metrics.CalculateForSeller(ctx, id)
	if err != nil {
		stdhttp.Error(w, err.Error(), stdhttp.StatusInternalServerError)
		return
	}

	writeJSON(w, metrics)
}
