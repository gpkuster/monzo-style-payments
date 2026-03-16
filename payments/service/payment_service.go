package service

import (
	"log"

	"github.com/google/uuid"
	"github.com/gpkuster/monzo-style-payments/payments/clients"
)

type PaymentService struct {
	store  *IdempotencyStore
	ledger *clients.LedgerClient
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		store:  NewIdempotencyStore(),
		ledger: clients.NewLedgerClient(),
	}
}

func (s *PaymentService) CreatePayment(idempotencyKey string, amount int64, currency string) (Payment, error) {
	log.Printf("CreatePayment: key=%s amount=%d currency=%s\n", idempotencyKey, amount, currency)

	// checks existing payment
	existing, ok := s.store.Get(idempotencyKey)

	if ok {
		log.Printf("CreatePayment: idempotency hit for key=%s\n", idempotencyKey)
		return existing, nil
	}

	payment := Payment{
		ID:       uuid.New().String(),
		Amount:   amount,
		Currency: currency,
		Status:   "created",
	}

	// stores idempotency key
	s.store.Set(idempotencyKey, payment)

	entries := []clients.Entry{
		{
			AccountID: "user_account",
			Amount:    -amount,
		},
		{
			AccountID: "merchant_account",
			Amount:    amount,
		},
	}

	// stores transaction in the ledger
	log.Printf("CreatePayment: calling ledger for payment %s\n", payment.ID)
	err := s.ledger.CreateTransaction(payment.ID, entries)
	if err != nil {
		log.Printf("CreatePayment: ledger call failed for payment %s: %v\n", payment.ID, err)
		return Payment{}, err
	}

	log.Printf("CreatePayment: ledger call successful for payment %s\n", payment.ID)
	return payment, nil
}
