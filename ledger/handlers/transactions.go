package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gpkuster/monzo-style-payments/ledger/service"
)

var ledgerService = service.NewLedgerService()

type CreateTransactionRequest struct {
	Reference string          `json:"reference"`
	Entries   []service.Entry `json:"entries"`
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	log.Println("Received CreateTransaction request")

	var req CreateTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("Error decoding CreateTransaction request: %v\n", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	tx, err := ledgerService.CreateTransaction(req.Reference, req.Entries)

	if err != nil {
		log.Printf("Error creating transaction in service: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Transaction created successfully: %s\n", tx.ID)
	json.NewEncoder(w).Encode(tx)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Printf("Received GetTransaction request for ID: %s\n", id)

	if id == "" {
		log.Println("GetTransaction: missing transaction ID")
		http.Error(w, "missing transaction ID", http.StatusBadRequest)
		return
	}

	tx, err := ledgerService.GetTransaction(id)

	if err != nil {
		log.Printf("GetTransaction: transaction %s not found\n", id)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Printf("Transaction %s retrieved successfully\n", id)
	json.NewEncoder(w).Encode(tx)
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	log.Println("Received GetAllTransactions request")

	txs := ledgerService.GetAllTransactions()

	log.Printf("Retrieved %d transactions\n", len(txs))
	json.NewEncoder(w).Encode(txs)
}
