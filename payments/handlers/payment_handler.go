package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gpkuster/monzo-style-payments/payments/service"
)

var paymentService = service.NewPaymentService()

type CreatePaymentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	log.Println("Received CreatePayment request")

	idempotencyKey := r.Header.Get("Idempotency-Key")

	if idempotencyKey == "" {
		log.Println("CreatePayment: missing Idempotency-Key header")
		http.Error(w, "missing Idempotency-Key header", http.StatusBadRequest)
		return
	}

	var req CreatePaymentRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding CreatePayment request: %v\n", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Processing payment for amount %d %s with idempotency key: %s\n", req.Amount, req.Currency, idempotencyKey)

	payment, err := paymentService.CreatePayment(
		idempotencyKey,
		req.Amount,
		req.Currency,
	)

	if err != nil {
		log.Printf("Error creating payment in service: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Payment processed successfully: %s\n", payment.ID)
	json.NewEncoder(w).Encode(payment)
}
