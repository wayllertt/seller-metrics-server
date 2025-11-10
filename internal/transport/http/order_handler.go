package transporthttp

import (
	"encoding/json"
	"net/http"

	"seller-metrics-server/internal/usecase"
)

type OrderHandler struct {
	orders *usecase.OrderService
}

func NewOrderHandler(o *usecase.OrderService) *OrderHandler {
	return &OrderHandler{orders: o}
}

// POST /orders/demo
func (h *OrderHandler) CreateDemoOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	if err := h.orders.CreateDemoOrders(ctx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{"status": "demo orders created"})
}

// POST /orders
type createOrderRequest struct {
	SellerID           int `json:"seller_id"`
	PromisedAfterHours int `json:"promised_after_hours"`
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req createOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if req.SellerID == 0 || req.PromisedAfterHours <= 0 {
		http.Error(w, "seller_id and promised_after_hours are required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := h.orders.CreateOrder(ctx, req.SellerID, req.PromisedAfterHours); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, map[string]string{"status": "order created"})
}
