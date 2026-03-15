package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gpkuster/monzo-style-payments/internal/service"
)

var paymentService = service.NewPaymentService()

type CreatePaymentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {

	idempotencyKey := r.Header.Get("Idempotency-Key")

	if idempotencyKey == "" {
		http.Error(w, "missing Idempotency-Key header", http.StatusBadRequest)
		return
	}

	var req CreatePaymentRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	payment := paymentService.CreatePayment(
		idempotencyKey,
		req.Amount,
		req.Currency,
	)

	json.NewEncoder(w).Encode(payment)
}
