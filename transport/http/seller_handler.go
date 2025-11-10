package http

import (
	"encoding/json"
	"net/http"

	"seller-metrics/internal/usecase"
)

type SellerHandler struct {
	sellerService *usecase.SellerService
}

func NewSellerHandler(s *usecase.SellerService) *SellerHandler {
	return &SellerHandler{sellerService: s}
}

func (h *SellerHandler) ListWithMetrics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := h.sellerService.ListWithMetrics(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, data)
}
