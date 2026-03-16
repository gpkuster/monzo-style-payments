package main

import (
	"log"
	"net/http"

	"github.com/gpkuster/monzo-style-payments/payments/handlers"
)

func main() {

	http.HandleFunc("/payments", handlers.CreatePayment)

	log.Println("payments-service running on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
