package transporthttp

import (
	stdhttp "net/http"

	"seller-metrics-server/usecase"
)

type SellerHandler struct {
	sellerService *usecase.SellerService
}

func NewSellerHandler(s *usecase.SellerService) *SellerHandler {
	return &SellerHandler{sellerService: s}
}

func (h *SellerHandler) ListWithMetrics(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()
	data, err := h.sellerService.ListWithMetrics(ctx)
	if err != nil {
		stdhttp.Error(w, err.Error(), stdhttp.StatusInternalServerError)
		return
	}
	writeJSON(w, data)
}
