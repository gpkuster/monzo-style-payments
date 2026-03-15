package handlers

import (
	"encoding/json"
	"net/http"
)

type CreatePaymentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type CreatePaymentResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {

	var req CreatePaymentRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	resp := CreatePaymentResponse{
		ID:     "payment_123",
		Status: "created",
	}

	json.NewEncoder(w).Encode(resp)
}
