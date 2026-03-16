package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gpkuster/monzo-style-payments/ledger/service"
)

var ledgerService = service.NewLedgerService()

type CreateTransactionRequest struct {
	Reference string          `json:"reference"`
	Entries   []service.Entry `json:"entries"`
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var req CreateTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	tx, err := ledgerService.CreateTransaction(req.Reference, req.Entries)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(tx)
}
