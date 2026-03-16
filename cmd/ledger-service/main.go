package main

import (
	"log"
	"net/http"

	"github.com/gpkuster/monzo-style-payments/ledger/handlers"
)

func main() {

	http.HandleFunc("/transactions", handlers.CreateTransaction)

	log.Println("ledger-service running on :8081")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
